package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/mail"
	"os"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/list"
	"github.com/pocketbase/pocketbase/tools/mailer"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/webhook"
)

func findTotalFilesByOwner(dao *daos.Dao, collection *models.Collection, ownerId string) (int, error) {
	var total int

	err := dao.RecordQuery(collection).
		Select("count(*)").
		AndWhere(dbx.HashExp{"owner": ownerId}).
		Row(&total)

	return total, err
}

func isUserSubscribed(dao *daos.Dao, authRecord *models.Record) bool {
	record, _ := dao.FindFirstRecordByData("subscriptions", "user", authRecord.Id)
	return record != nil
}

func getFilePermission(app *pocketbase.PocketBase, fileRecord *models.Record, authRecord *models.Record) (string, error) {

	// file is owned
	if fileRecord.GetString("owner") == authRecord.Id {
		return "edit", nil
	}

	// file team is owned
	if fileRecord.GetString("team") != "" {
		team, err := app.Dao().FindRecordById("teams", fileRecord.GetString("team"))

		if err != nil {
			return "", err
		}

		if team.GetString("owner") == authRecord.Id {
			return "edit", nil
		}
	}

	// file is shared -> check file share permissions
	fileShares, err := app.Dao().FindRecordsByExpr("file_share", dbx.HashExp{"id": list.ToInterfaceSlice(fileRecord.GetStringSlice("collaborators")), "user": authRecord.Id})
	if err != nil {
		return "", err
	}
	if len(fileShares) == 1 {
		return fileShares[0].GetString("permission"), nil
	}

	// file is in team -> check team member permissions
	if fileRecord.GetString("team") != "" {
		teamMembers, err := app.Dao().FindRecordsByExpr("team_members", dbx.HashExp{"team": fileRecord.GetString("team"), "user": authRecord.Id})

		if err != nil {
			return "", err
		}

		if len(teamMembers) == 1 {
			return teamMembers[0].GetString("permission"), nil
		}
	}

	return "view", nil
}

func main() {
	stripe.Key = os.Getenv("STRIPE_API_KEY")

	app := pocketbase.New()

	app.OnRecordsListRequest("files").Add(func(e *core.RecordsListEvent) error {
		authRecord, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)

		if authRecord == nil {
			return nil
		}

		for _, r := range e.Records {
			permission, err := getFilePermission(app, r, authRecord)
			if err != nil {
				return err
			}
			r.SetExpand(map[string]any{"permission": map[string]any{"value": permission}})
		}

		return nil
	})

	app.OnRecordViewRequest("files").Add(func(e *core.RecordViewEvent) error {
		authRecord, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)

		if authRecord == nil {
			return nil
		}

		permission, err := getFilePermission(app, e.Record, authRecord)
		if err != nil {
			return err
		}
		e.Record.SetExpand(map[string]any{"permission": map[string]any{"value": permission}})

		return nil
	})

	// check whether a default user has reached the score/file limit
	app.OnRecordBeforeCreateRequest("files", "scores").Add(func(e *core.RecordCreateEvent) error {
		authRecord, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)

		if admin, _ := e.HttpContext.Get(apis.ContextAdminKey).(*models.Admin); admin != nil {
			return nil
		}

		total, err := findTotalFilesByOwner(app.Dao(), e.Record.Collection(), e.Record.GetString("owner"))
		if err != nil || (!isUserSubscribed(app.Dao(), authRecord) && total >= 2) {
			return apis.NewBadRequestError(fmt.Sprintf("Cannot create more %s!", e.Record.Collection().Name), err)
		}

		return nil
	})

	app.OnRecordAfterCreateRequest("team_members").Add(func(e *core.RecordCreateEvent) error {

		team, err := app.Dao().FindRecordById("teams", e.Record.GetString("team"))

		if err != nil {
			return err
		}

		owner, err := app.Dao().FindRecordById("users", team.GetString("owner"))

		if err != nil {
			return err
		}

		user, err := app.Dao().FindRecordById("users", e.Record.GetString("user"))

		if err != nil {
			return err
		}

		link := fmt.Sprintf(os.Getenv("APP_URL")+"/dashboard/teams/accept?team_id=%s", e.Record.Id)
		subject := fmt.Sprintf("%s has invited you to join %s", owner.GetString("username"), team.GetString("name"))
		message := &mailer.Message{
			From: mail.Address{
				Address: app.Settings().Meta.SenderAddress,
				Name:    app.Settings().Meta.SenderName,
			},
			To:      []mail.Address{{Address: user.Email()}},
			Subject: subject,
			HTML: fmt.Sprintf(`
			<div style="max-width: 650px; margin: auto"><svg width="128" viewBox="0 0 464 190" fill="none"
			xmlns="http://www.w3.org/2000/svg">
			<path d="M122.156 68L80.2031 188H42.7031L0.75 68H35.9062L60.8281 153.859H62.0781L86.9219 68H122.156Z"
				fill="black" />
			<path d="M123.787 188V68H157.069V188H123.787Z" fill="black" />
			<path
				d="M168.329 188V68H200.048V89.1719H201.454C203.954 82.1406 208.121 76.5937 213.954 72.5312C219.787 68.4687 226.766 66.4375 234.891 66.4375C243.121 66.4375 250.126 68.4948 255.907 72.6094C261.688 76.6719 265.542 82.1927 267.47 89.1719H268.72C271.167 82.2969 275.595 76.8021 282.001 72.6875C288.459 68.5208 296.089 66.4375 304.891 66.4375C316.089 66.4375 325.178 70.0052 332.157 77.1406C339.188 84.224 342.704 94.276 342.704 107.297V188H309.501V113.859C309.501 107.193 307.73 102.193 304.188 98.8594C300.647 95.526 296.22 93.8594 290.907 93.8594C284.865 93.8594 280.152 95.7865 276.766 99.6406C273.381 103.443 271.688 108.469 271.688 114.719V188H239.423V113.156C239.423 107.271 237.73 102.583 234.345 99.0937C231.011 95.6042 226.61 93.8594 221.141 93.8594C217.443 93.8594 214.11 94.7969 211.141 96.6719C208.225 98.4948 205.907 101.073 204.188 104.406C202.47 107.688 201.61 111.542 201.61 115.969V188H168.329Z"
				fill="black" />
			<path
				d="M430.546 136.906V68H463.827V188H431.874V166.203H430.624C427.916 173.234 423.411 178.885 417.109 183.156C410.859 187.427 403.228 189.563 394.218 189.563C386.197 189.563 379.14 187.74 373.046 184.094C366.952 180.448 362.187 175.266 358.749 168.547C355.364 161.828 353.645 153.781 353.593 144.406V68H386.874V138.469C386.926 145.552 388.827 151.151 392.577 155.266C396.327 159.38 401.353 161.437 407.655 161.437C411.666 161.437 415.416 160.526 418.905 158.703C422.395 156.828 425.208 154.068 427.343 150.422C429.53 146.776 430.598 142.271 430.546 136.906Z"
				fill="black" />
			<path
				d="M174.107 46.7572C171.253 46.7572 168.719 46.2816 166.506 45.3303C164.313 44.3597 162.575 42.9037 161.294 40.9624C160.032 39.0017 159.401 36.546 159.401 33.5953C159.401 31.1104 159.838 29.0138 160.712 27.3055C161.585 25.5972 162.789 24.2091 164.322 23.1414C165.856 22.0737 167.623 21.2681 169.622 20.7245C171.622 20.1615 173.757 19.783 176.028 19.5889C178.571 19.3559 180.62 19.1133 182.173 18.8609C183.726 18.5891 184.852 18.2106 185.55 17.7252C186.269 17.2205 186.628 16.5119 186.628 15.5995V15.4539C186.628 13.9591 186.113 12.8041 185.084 11.9887C184.056 11.1734 182.668 10.7657 180.92 10.7657C179.037 10.7657 177.523 11.1734 176.378 11.9887C175.232 12.8041 174.504 13.93 174.194 15.3666L161.061 14.9007C161.449 12.1829 162.449 9.75624 164.06 7.62083C165.691 5.466 167.924 3.77708 170.758 2.55406C173.611 1.31164 177.038 0.69043 181.037 0.69043C183.891 0.69043 186.521 1.03015 188.928 1.7096C191.335 2.36964 193.432 3.34029 195.218 4.62154C197.004 5.88337 198.382 7.4364 199.353 9.28063C200.343 11.1249 200.838 13.2311 200.838 15.5995V46.0001H187.443V39.7686H187.094C186.298 41.2828 185.279 42.564 184.036 43.6123C182.813 44.6606 181.367 45.4468 179.697 45.971C178.047 46.4951 176.184 46.7572 174.107 46.7572ZM178.504 37.439C180.037 37.439 181.415 37.1284 182.638 36.5072C183.881 35.886 184.871 35.0318 185.609 33.9447C186.346 32.8382 186.715 31.5569 186.715 30.1009V25.8495C186.307 26.0631 185.812 26.2572 185.23 26.4319C184.667 26.6066 184.046 26.7716 183.366 26.9269C182.687 27.0822 181.988 27.2181 181.27 27.3346C180.552 27.4511 179.862 27.5579 179.202 27.6549C177.863 27.8685 176.718 28.1985 175.766 28.645C174.834 29.0915 174.116 29.6739 173.611 30.3921C173.126 31.091 172.883 31.9258 172.883 32.8964C172.883 34.3718 173.408 35.4977 174.456 36.2742C175.524 37.0508 176.873 37.439 178.504 37.439Z"
				fill="black" />
			<path
				d="M209.545 62.7728V1.27281H223.638V8.93119H224.075C224.658 7.57229 225.483 6.26192 226.55 5.00009C227.638 3.73825 229.016 2.70937 230.685 1.91344C232.374 1.0981 234.393 0.69043 236.742 0.69043C239.848 0.69043 242.75 1.50577 245.449 3.13645C248.167 4.76713 250.36 7.2811 252.03 10.6784C253.699 14.0756 254.534 18.4047 254.534 23.6656C254.534 28.7323 253.728 32.974 252.117 36.3907C250.525 39.8074 248.37 42.3699 245.653 44.0782C242.954 45.7865 239.955 46.6407 236.655 46.6407C234.403 46.6407 232.452 46.2719 230.802 45.5342C229.152 44.7965 227.764 43.8258 226.638 42.6222C225.531 41.4186 224.677 40.1277 224.075 38.7494H223.784V62.7728H209.545ZM223.493 23.6365C223.493 26.0436 223.813 28.1402 224.454 29.9262C225.114 31.7122 226.055 33.1002 227.278 34.0903C228.521 35.0609 230.006 35.5463 231.734 35.5463C233.481 35.5463 234.966 35.0609 236.189 34.0903C237.412 33.1002 238.334 31.7122 238.955 29.9262C239.596 28.1402 239.916 26.0436 239.916 23.6365C239.916 21.2293 239.596 19.1424 238.955 17.3758C238.334 15.6092 237.412 14.2406 236.189 13.27C234.985 12.2993 233.5 11.814 231.734 11.814C229.986 11.814 228.501 12.2896 227.278 13.2409C226.055 14.1921 225.114 15.551 224.454 17.3176C223.813 19.0841 223.493 21.1904 223.493 23.6365Z"
				fill="black" />
			<path
				d="M261.996 62.7728V1.27281H276.09V8.93119H276.526C277.109 7.57229 277.934 6.26192 279.002 5.00009C280.089 3.73825 281.467 2.70937 283.137 1.91344C284.825 1.0981 286.844 0.69043 289.193 0.69043C292.299 0.69043 295.202 1.50577 297.9 3.13645C300.618 4.76713 302.811 7.2811 304.481 10.6784C306.15 14.0756 306.985 18.4047 306.985 23.6656C306.985 28.7323 306.18 32.974 304.568 36.3907C302.976 39.8074 300.822 42.3699 298.104 44.0782C295.405 45.7865 292.406 46.6407 289.106 46.6407C286.854 46.6407 284.903 46.2719 283.253 45.5342C281.603 44.7965 280.215 43.8258 279.089 42.6222C277.982 41.4186 277.128 40.1277 276.526 38.7494H276.235V62.7728H261.996ZM275.944 23.6365C275.944 26.0436 276.264 28.1402 276.905 29.9262C277.565 31.7122 278.507 33.1002 279.73 34.0903C280.972 35.0609 282.457 35.5463 284.185 35.5463C285.932 35.5463 287.417 35.0609 288.64 34.0903C289.863 33.1002 290.785 31.7122 291.406 29.9262C292.047 28.1402 292.367 26.0436 292.367 23.6365C292.367 21.2293 292.047 19.1424 291.406 17.3758C290.785 15.6092 289.863 14.2406 288.64 13.27C287.436 12.2993 285.951 11.814 284.185 11.814C282.438 11.814 280.953 12.2896 279.73 13.2409C278.507 14.1921 277.565 15.551 276.905 17.3176C276.264 19.0841 275.944 21.1904 275.944 23.6365Z"
				fill="black" />
			<circle cx="140.5" cy="39.5" r="11.5" fill="black" />
		</svg>
	
	
	
	
		<div style="text-align: center;">
			<div style="padding: 64px 32px">
				<p
					style="font-size:24px;line-height:32px;font-weight:normal;font-family:'Inter',Arial,Helvetica,sans-serif;color:#000;font-weight:normal;margin:0;padding-bottom:0px">
					<strong>%s</strong> has invited you to join the team <strong>%s</strong>
				</p>
			</div>
			<div>
				<p style="font-size:20px;line-height:28px;font-weight:normal;font-family:'Inter',Arial,Helvetica,sans-serif;color:#000;font-weight:normal;margin:0;padding-bottom:0px">
					<a href="%s" target="_blank" rel="noopener" style="outline: 2px solid black; border-radius: 10px; padding: 24px; text-decoration: none; color:#000">Accept invite</a>
				</p>
			</div>
		</div>	
			`, owner.GetString("username"), team.GetString("name"), link),
		}

		return app.NewMailClient().Send(message)
	})

	app.OnRealtimeAfterSubscribeRequest().Add(func(e *core.RealtimeSubscribeEvent) error {
		authRecord, _ := e.Client.Get(apis.ContextAuthRecordKey).(*models.Record)

		if authRecord != nil && len(e.Subscriptions) > 0 {
			fileData, err := app.Dao().FindRecordById("file_data", strings.Split(e.Subscriptions[0], "/")[1])
			if err != nil {
				return err
			}

			existingSubscriptions, err := app.Dao().FindRecordsByExpr("file_data_editors", dbx.HashExp{"file_data": fileData.Id, "user": authRecord.Id})

			if err != nil {
				return err
			}

			if len(existingSubscriptions) > 0 {
				sub := existingSubscriptions[0]
				sub.Set("subscription_id", e.Client.Id())

				if err := app.Dao().SaveRecord(sub); err != nil {
					return err
				}

			} else {
				collection, err := app.Dao().FindCollectionByNameOrId("file_data_editors")
				if err != nil {
					return err
				}

				fileDataEditor := models.NewRecord(collection)
				fileDataEditor.Set("file_data", fileData.Id)
				fileDataEditor.Set("user", authRecord.Id)
				fileDataEditor.Set("subscription_id", e.Client.Id())

				if err := app.Dao().SaveRecord(fileDataEditor); err != nil {
					return err
				}

				editors := append(fileData.GetStringSlice("editors"), fileDataEditor.Id)
				fileData.Set("editors", editors)

				if err := app.Dao().SaveRecord(fileData); err != nil {
					return err
				}
			}
		}

		return nil
	})

	app.OnRealtimeDisconnectRequest().Add(func(e *core.RealtimeDisconnectEvent) error {
		authRecord, _ := e.Client.Get(apis.ContextAuthRecordKey).(*models.Record)

		if authRecord != nil {
			fileDataEditor, err := app.Dao().FindFirstRecordByData("file_data_editors", "subscription_id", e.Client.Id())
			if err != nil {
				return err
			}

			if err != nil {
				return err
			}

			if err := app.Dao().DeleteRecord(fileDataEditor); err != nil {
				return err
			}
		}
		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodPost,
			Path:   "/api/stripe/webhook",
			Handler: func(c echo.Context) error {
				collection, err := app.Dao().FindCollectionByNameOrId("subscriptions")
				if err != nil {
					return err
				}
				request := c.Request()

				payload, err := io.ReadAll(request.Body)
				if err != nil {
					return apis.NewBadRequestError("Error reading request body", err)
				}

				endpointSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")

				// Pass the request body and Stripe-Signature header to ConstructEvent, along
				// with the webhook signing key.
				event, err := webhook.ConstructEvent(payload, request.Header.Get("Stripe-Signature"),
					endpointSecret)

				if err != nil {
					return apis.NewBadRequestError("Error verifying webhook signature", err)
				}

				dataObject := event.Data.Object

				// Unmarshal the event data into an appropriate struct depending on its Type
				switch event.Type {
				case "checkout.session.completed":
					userId := dataObject["client_reference_id"]
					stripeCustomerId := dataObject["customer"]
					stripeSubscriptionId := dataObject["subscription"]

					record := models.NewRecord(collection)
					record.Set("stripe_subscription_id", stripeSubscriptionId)
					record.Set("stripe_customer_id", stripeCustomerId)
					record.Set("user", userId)
					record.Set("status", "active")

					if err := app.Dao().SaveRecord(record); err != nil {
						return err
					}

				case "invoice.paid":
					stripeSubscriptionId := dataObject["stripe_id"]
					record, err := app.Dao().FindFirstRecordByData("subscriptions", "stripe_subscription_id", stripeSubscriptionId)
					if err != nil {
						return err
					}

					record.Set("status", "active")

					if err := app.Dao().SaveRecord(record); err != nil {
						return err
					}

				case "invoice.payment_failed":
					stripeSubscriptionId := dataObject["stripe_id"]
					record, err := app.Dao().FindFirstRecordByData("subscriptions", "stripe_subscription_id", stripeSubscriptionId)
					if err != nil {
						return err
					}

					record.Set("status", "unpaid")

					if err := app.Dao().SaveRecord(record); err != nil {
						return err
					}
				case "customer.subscription.updated":
					stripeSubscriptionId := dataObject["stripe_id"]
					record, err := app.Dao().FindFirstRecordByData("subscriptions", "stripe_subscription_id", stripeSubscriptionId)
					if err != nil {
						return err
					}
					status := dataObject["status"]
					record.Set("status", status)

					if err := app.Dao().SaveRecord(record); err != nil {
						return err
					}
				case "customer.subscription.deleted":
					stripeSubscriptionId := dataObject["stripe_id"]
					record, err := app.Dao().FindFirstRecordByData("subscriptions", "stripe_subscription_id", stripeSubscriptionId)
					if err != nil {
						return err
					}

					if err := app.Dao().DeleteRecord(record); err != nil {
						return err
					}
				default:
					fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
				}

				return c.JSON(http.StatusOK, nil)
			},
		})
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
