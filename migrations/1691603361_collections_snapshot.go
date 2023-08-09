package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "_pb_users_auth_",
				"created": "2023-01-29 09:56:42.813Z",
				"updated": "2023-07-25 17:41:13.856Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "z4bn5gxn",
						"name": "avatar",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": 7,
							"pattern": ""
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `__pb_users_auth__created_idx` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `created` + "`" + `)"
				],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "",
				"updateRule": "id = @request.auth.id",
				"deleteRule": "id = @request.auth.id",
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": true,
					"allowUsernameAuth": true,
					"exceptEmailDomains": null,
					"manageRule": null,
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"requireEmail": true
				}
			},
			{
				"id": "1b8m8nuntgbyzal",
				"created": "2023-01-29 09:57:27.358Z",
				"updated": "2023-07-25 17:41:13.859Z",
				"name": "scores",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "brrrbvl0",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "m3lln0kx",
						"name": "data",
						"type": "file",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [],
							"thumbs": [],
							"protected": false
						}
					},
					{
						"system": false,
						"id": "tpo41d1x",
						"name": "thumbnail",
						"type": "file",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/svg+xml"
							],
							"thumbs": [],
							"protected": false
						}
					},
					{
						"system": false,
						"id": "bixobfsh",
						"name": "public",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "p7qtcgmp",
						"name": "meta",
						"type": "relation",
						"required": false,
						"unique": true,
						"options": {
							"collectionId": "6dqgglubeh5alx5",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "34wtunfq",
						"name": "owner",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `_1b8m8nuntgbyzal_created_idx` + "`" + ` ON ` + "`" + `scores` + "`" + ` (` + "`" + `created` + "`" + `)",
					"CREATE UNIQUE INDEX \"idx_unique_p7qtcgmp\" on \"scores\" (\"meta\")"
				],
				"listRule": "@request.auth.id != \"\" && (owner = @request.auth.id || public = true)",
				"viewRule": "@request.auth.id != \"\" && (owner = @request.auth.id || public = true)",
				"createRule": "@request.auth.id != \"\" && owner = @request.auth.id",
				"updateRule": "@request.auth.id != \"\" && (owner = @request.auth.id)",
				"deleteRule": "@request.auth.id != \"\" && (owner = @request.auth.id)",
				"options": {}
			},
			{
				"id": "6dqgglubeh5alx5",
				"created": "2023-01-29 09:57:27.358Z",
				"updated": "2023-07-25 17:41:13.860Z",
				"name": "score_meta",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "obrjacru",
						"name": "composer",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "4xk4oczy",
						"name": "date",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "81ui4tbf",
						"name": "instruments",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "oribv0gb",
						"name": "keys",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "zu4qm7uz",
						"name": "language",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "4rxzz09s",
						"name": "opus",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "t6l3hkbo",
						"name": "times",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "wfypwsyb",
						"name": "lyrics",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `_6dqgglubeh5alx5_created_idx` + "`" + ` ON ` + "`" + `score_meta` + "`" + ` (` + "`" + `created` + "`" + `)"
				],
				"listRule": "@request.auth.id != \"\" && @collection.scores.meta = id && (@collection.scores.owner = @request.auth.id || @collection.scores.public = true)",
				"viewRule": "@request.auth.id != \"\" && @collection.scores.meta ?= id && (@collection.scores.owner ?= @request.auth.id || @collection.scores.public ?= true)",
				"createRule": "@request.auth.id != \"\"",
				"updateRule": "@request.auth.id != \"\" && @collection.scores.meta ?= id && (@collection.scores.owner ?= @request.auth.id)",
				"deleteRule": "@request.auth.id != \"\" && @collection.scores.meta ?= id && (@collection.scores.owner ?= @request.auth.id)",
				"options": {}
			},
			{
				"id": "2xrpwt0zs5kyuel",
				"created": "2023-01-29 09:57:27.358Z",
				"updated": "2023-07-25 17:41:13.862Z",
				"name": "files",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "q8as8je9",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "an4uo2nc",
						"name": "public",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "jirmgif8",
						"name": "owner",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "736ddpbx",
						"name": "collaborators",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"collectionId": "v6gkd5kd64fmyuv",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "kayabqel",
						"name": "team",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"collectionId": "nrmy6xhnaygd60n",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `_2xrpwt0zs5kyuel_created_idx` + "`" + ` ON ` + "`" + `files` + "`" + ` (` + "`" + `created` + "`" + `)"
				],
				"listRule": "@request.auth.id != \"\" && (owner = @request.auth.id || collaborators.user ?= @request.auth.id || \n team.owner = @request.auth.id || (@collection.team_members.team = team && @collection.team_members.user = @request.auth.id) || public = true)",
				"viewRule": "@request.auth.id != \"\" && (owner = @request.auth.id || collaborators.user ?= @request.auth.id || \n team.owner = @request.auth.id ||  \n  (@collection.team_members.team = team && @collection.team_members.user = @request.auth.id) || public = true)",
				"createRule": "@request.auth.id != \"\" && (@request.data.owner = @request.auth.id)",
				"updateRule": "@request.auth.id != \"\" && (owner = @request.auth.id || \n team.owner = @request.auth.id ||  (@collection.team_members.team = team && @collection.team_members.user = @request.auth.id && @collection.team_members.permission = \"edit\"))",
				"deleteRule": "@request.auth.id != \"\" && (owner = @request.auth.id|| \n team.owner = @request.auth.id ||  (@collection.team_members.team = team && @collection.team_members.user = @request.auth.id && @collection.team_members.permission = \"edit\"))",
				"options": {}
			},
			{
				"id": "bfgxjojefgb666z",
				"created": "2023-01-29 09:57:27.358Z",
				"updated": "2023-07-25 17:41:13.866Z",
				"name": "file_data",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "chi0oxgn",
						"name": "json",
						"type": "json",
						"required": true,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "duhosiw1",
						"name": "file",
						"type": "relation",
						"required": true,
						"unique": true,
						"options": {
							"collectionId": "2xrpwt0zs5kyuel",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "kkfic81n",
						"name": "editors",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"collectionId": "8elpvo708zjue5o",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": null
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `_bfgxjojefgb666z_created_idx` + "`" + ` ON ` + "`" + `file_data` + "`" + ` (` + "`" + `created` + "`" + `)",
					"CREATE UNIQUE INDEX \"idx_unique_duhosiw1\" on \"file_data\" (\"file\")"
				],
				"listRule": "@request.auth.id != \"\" && file.owner = @request.auth.id",
				"viewRule": "file.owner = @request.auth.id || file.collaborators.user.id ?= @request.auth.id || \n file.team.owner = @request.auth.id || (@collection.team_members.team = file.team && @collection.team_members.user = @request.auth.id) || file.public = true",
				"createRule": "@request.auth.id != \"\" && file.owner = @request.auth.id",
				"updateRule": "@request.auth.id != \"\" && (file.owner = @request.auth.id || \n file.team.owner = @request.auth.id || (file.collaborators.user.id ?= @request.auth.id && file.collaborators.permission = \"edit\") ||\n(@collection.team_members.team = file.team && @collection.team_members.user = @request.auth.id && @collection.team_members.permission = \"edit\"))",
				"deleteRule": "@request.auth.id != \"\" && file.owner = @request.auth.id",
				"options": {}
			},
			{
				"id": "v6gkd5kd64fmyuv",
				"created": "2023-01-29 09:57:27.358Z",
				"updated": "2023-07-25 17:41:13.871Z",
				"name": "file_share",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "n0s6mab7",
						"name": "permission",
						"type": "select",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"view",
								"edit"
							]
						}
					},
					{
						"system": false,
						"id": "mszovgvi",
						"name": "user",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `_v6gkd5kd64fmyuv_created_idx` + "`" + ` ON ` + "`" + `file_share` + "`" + ` (` + "`" + `created` + "`" + `)"
				],
				"listRule": "@request.auth.id != \"\" && (@collection.files.collaborators.id ?= id && (@collection.files.owner ?= @request.auth.id || user = @request.auth.id))",
				"viewRule": "@request.auth.id != \"\" && (@collection.files.collaborators.id ?= id && (@collection.files.owner ?= @request.auth.id || user = @request.auth.id))",
				"createRule": "@request.auth.id != \"\" && @collection.subscriptions.user ?= @request.auth.id",
				"updateRule": "@request.auth.id != \"\" && (@collection.files.collaborators.id ?= id && @collection.files.owner ?= @request.auth.id) && @collection.subscriptions.user = @request.auth.id",
				"deleteRule": "@request.auth.id != \"\" && (@collection.files.collaborators.id ?= id && @collection.files.owner ?= @request.auth.id) && @collection.subscriptions.user = @request.auth.id",
				"options": {}
			},
			{
				"id": "tqh82amh27108kk",
				"created": "2023-01-29 09:57:27.358Z",
				"updated": "2023-07-25 17:41:13.875Z",
				"name": "subscriptions",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "wejd2o0q",
						"name": "stripe_subscription_id",
						"type": "text",
						"required": true,
						"unique": true,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "mpxd9cuk",
						"name": "stripe_customer_id",
						"type": "text",
						"required": true,
						"unique": true,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "k6ukkwzb",
						"name": "user",
						"type": "relation",
						"required": true,
						"unique": true,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "beu25egs",
						"name": "status",
						"type": "select",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"active",
								"past_due",
								"unpaid",
								"canceled",
								"incomplete",
								"incomplete_expired",
								"trialing"
							]
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `_tqh82amh27108kk_created_idx` + "`" + ` ON ` + "`" + `subscriptions` + "`" + ` (` + "`" + `created` + "`" + `)",
					"CREATE UNIQUE INDEX \"idx_unique_wejd2o0q\" on \"subscriptions\" (\"stripe_subscription_id\")",
					"CREATE UNIQUE INDEX \"idx_unique_mpxd9cuk\" on \"subscriptions\" (\"stripe_customer_id\")",
					"CREATE UNIQUE INDEX \"idx_unique_k6ukkwzb\" on \"subscriptions\" (\"user\")"
				],
				"listRule": "@request.auth.id != \"\" && user = @request.auth.id",
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "4xryl53w59yf60r",
				"created": "2023-01-29 09:57:27.358Z",
				"updated": "2023-07-25 17:41:13.880Z",
				"name": "email_settings",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "gplgb9it",
						"name": "share",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "fileho6h",
						"name": "team",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "kfd6flzq",
						"name": "changelog",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "tnzwgf0l",
						"name": "marketing",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "ngkeik6n",
						"name": "user",
						"type": "relation",
						"required": true,
						"unique": true,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `_4xryl53w59yf60r_created_idx` + "`" + ` ON ` + "`" + `email_settings` + "`" + ` (` + "`" + `created` + "`" + `)",
					"CREATE UNIQUE INDEX \"idx_unique_ngkeik6n\" on \"email_settings\" (\"user\")"
				],
				"listRule": "@request.auth.id != \"\" && user = @request.auth.id",
				"viewRule": null,
				"createRule": "@request.auth.id != \"\" && @request.data.user = @request.auth.id",
				"updateRule": "@request.auth.id != \"\" && user = @request.auth.id",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "57e7nnymic6xocc",
				"created": "2023-01-29 09:57:27.358Z",
				"updated": "2023-07-25 17:41:13.882Z",
				"name": "editor_settings",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "cdgpvxnp",
						"name": "score",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "hpccaheo",
						"name": "plot",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "ohdclhzk",
						"name": "minimap",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "np2fxdaw",
						"name": "pixel_grid",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "3a4r1f0o",
						"name": "grid_columns",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "zvd6la8d",
						"name": "grid_rows",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "3pblqibc",
						"name": "tutorial_completed",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "77bfob6c",
						"name": "user",
						"type": "relation",
						"required": true,
						"unique": true,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "gvtzqjkr",
						"name": "player_volume",
						"type": "number",
						"required": false,
						"unique": false,
						"options": {
							"min": 0,
							"max": 100
						}
					},
					{
						"system": false,
						"id": "ghgkikvd",
						"name": "player_tempo",
						"type": "number",
						"required": true,
						"unique": false,
						"options": {
							"min": 20,
							"max": 240
						}
					},
					{
						"system": false,
						"id": "mgeggtsk",
						"name": "display_show_title",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "t5rwuqja",
						"name": "display_show_composer",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "fosdbpep",
						"name": "display_show_lyrics",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "ticnxzj3",
						"name": "display_show_measure_numbers",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "2dfphp5c",
						"name": "display_follow_cursor",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "pjeaefsk",
						"name": "display_show_part_names",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `_57e7nnymic6xocc_created_idx` + "`" + ` ON ` + "`" + `editor_settings` + "`" + ` (` + "`" + `created` + "`" + `)",
					"CREATE UNIQUE INDEX \"idx_unique_77bfob6c\" on \"editor_settings\" (\"user\")"
				],
				"listRule": "@request.auth.id != \"\" && user = @request.auth.id",
				"viewRule": null,
				"createRule": "@request.auth.id != \"\" && user = @request.auth.id",
				"updateRule": "@request.auth.id != \"\" && user = @request.auth.id",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "8elpvo708zjue5o",
				"created": "2023-01-29 10:42:25.503Z",
				"updated": "2023-07-25 17:41:13.884Z",
				"name": "file_data_editors",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "sdgjd51w",
						"name": "file_data",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "bfgxjojefgb666z",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "wqkixi6e",
						"name": "user",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "nkem7mkd",
						"name": "subscription_id",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `_8elpvo708zjue5o_created_idx` + "`" + ` ON \"file_data_editors\" (` + "`" + `created` + "`" + `)"
				],
				"listRule": "@request.auth.id != \"\" && (file_data.file.owner = @request.auth.id || file_data.file.collaborators.user.id ?= @request.auth.id || file_data.file.team.owner = @request.auth.id || (@collection.team_members.team = file_data.file.team && @collection.team_members.user = @request.auth.id))",
				"viewRule": "@request.auth.id != \"\" && (file_data.file.owner = @request.auth.id || file_data.file.collaborators.user.id ?= @request.auth.id || file_data.file.team.owner = @request.auth.id || (@collection.team_members.team = file_data.file.team && @collection.team_members.user = @request.auth.id))",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "nrmy6xhnaygd60n",
				"created": "2023-02-18 09:49:00.105Z",
				"updated": "2023-07-25 17:41:13.886Z",
				"name": "teams",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "ru3bbk8w",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "krvzxchb",
						"name": "icon",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/png",
								"image/jpeg",
								"image/svg+xml"
							],
							"thumbs": [],
							"protected": false
						}
					},
					{
						"system": false,
						"id": "jyufjk4j",
						"name": "owner",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `_nrmy6xhnaygd60n_created_idx` + "`" + ` ON ` + "`" + `teams` + "`" + ` (` + "`" + `created` + "`" + `)"
				],
				"listRule": "@request.auth.id != \"\" && (owner = @request.auth.id || (@collection.team_members.team = id && @collection.team_members.user = @request.auth.id && @collection.team_members.status = \"active\"))",
				"viewRule": "@request.auth.id != \"\" && (owner = @request.auth.id || (@collection.team_members.team = id && @collection.team_members.user = @request.auth.id && @collection.team_members.status = \"active\"))",
				"createRule": "@request.auth.id != \"\" && @collection.subscriptions.user ?= @request.auth.id",
				"updateRule": "@request.auth.id != \"\" && @collection.subscriptions.user ?= @request.auth.id && owner = @request.auth.id",
				"deleteRule": "@request.auth.id != \"\" && @collection.subscriptions.user ?= @request.auth.id && owner = @request.auth.id",
				"options": {}
			},
			{
				"id": "z79hqmr1d6axug2",
				"created": "2023-02-18 09:49:00.105Z",
				"updated": "2023-07-25 17:41:13.890Z",
				"name": "team_members",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "iary9ewd",
						"name": "team",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "nrmy6xhnaygd60n",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "pu8plxez",
						"name": "user",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "9rowcerb",
						"name": "permission",
						"type": "select",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"view",
								"edit"
							]
						}
					},
					{
						"system": false,
						"id": "nw9ikmmd",
						"name": "status",
						"type": "select",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"pending",
								"active"
							]
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `_z79hqmr1d6axug2_created_idx` + "`" + ` ON ` + "`" + `team_members` + "`" + ` (` + "`" + `created` + "`" + `)"
				],
				"listRule": "@request.auth.id != \"\" && (team.owner = @request.auth.id || user = @request.auth.id)",
				"viewRule": "@request.auth.id != \"\" &&  (team.owner = @request.auth.id || user = @request.auth.id)",
				"createRule": "@request.auth.id != \"\" && @collection.subscriptions.user ?= @request.auth.id && team.owner = @request.auth.id",
				"updateRule": "@request.auth.id != \"\" && ((user.id = @request.auth.id && \n@request.data.status = \"active\" && @request.data.permission:isset = false) || team.owner = @request.auth.id)",
				"deleteRule": "@request.auth.id != \"\" && (team.owner = @request.auth.id || user = @request.auth.id)",
				"options": {}
			},
			{
				"id": "2y3culndctnplgp",
				"created": "2023-02-19 11:57:03.607Z",
				"updated": "2023-07-25 17:41:13.891Z",
				"name": "file_favorites",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "75jfhncm",
						"name": "file",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "2xrpwt0zs5kyuel",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": [
								"id"
							]
						}
					},
					{
						"system": false,
						"id": "g076dsnx",
						"name": "user",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": [
								"id"
							]
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `_2y3culndctnplgp_created_idx` + "`" + ` ON ` + "`" + `file_favorites` + "`" + ` (` + "`" + `created` + "`" + `)"
				],
				"listRule": "@request.auth.id != \"\" && user = @request.auth.id",
				"viewRule": "@request.auth.id != \"\" && user = @request.auth.id",
				"createRule": "@request.auth.id != \"\" && @request.data.user = @request.auth.id",
				"updateRule": null,
				"deleteRule": "@request.auth.id != \"\" && user = @request.auth.id",
				"options": {}
			},
			{
				"id": "qschi6gclkglton",
				"created": "2023-07-31 09:35:28.127Z",
				"updated": "2023-08-08 19:03:22.894Z",
				"name": "plugins",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "fx2g4tuv",
						"name": "name",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "vmcobxwg",
						"name": "description",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 400,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "7vqtbnmv",
						"name": "config",
						"type": "json",
						"required": true,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "0vs65yj2",
						"name": "code",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "edongscw",
						"name": "public",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "60gd32ev",
						"name": "owner",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\" && (owner = @request.auth.id || public = true)",
				"viewRule": "@request.auth.id != \"\" && (owner = @request.auth.id || public = true)",
				"createRule": "@request.auth.id != \"\" && owner = @request.auth.id",
				"updateRule": "@request.auth.id != \"\" && owner = @request.auth.id",
				"deleteRule": "@request.auth.id != \"\" && owner = @request.auth.id",
				"options": {}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
