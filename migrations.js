migrate((db) => {
  const snapshot = [
    {
      "id": "_pb_users_auth_",
      "name": "users",
      "type": "auth",
      "system": false,
      "schema": [
        {
          "id": "z4bn5gxn",
          "name": "avatar",
          "type": "text",
          "system": false,
          "required": true,
          "unique": false,
          "options": {
            "min": null,
            "max": 7,
            "pattern": ""
          }
        }
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
      "name": "scores",
      "type": "base",
      "system": false,
      "schema": [
        {
          "id": "brrrbvl0",
          "name": "name",
          "type": "text",
          "system": false,
          "required": true,
          "unique": false,
          "options": {
            "min": null,
            "max": null,
            "pattern": ""
          }
        },
        {
          "id": "m3lln0kx",
          "name": "data",
          "type": "file",
          "system": false,
          "required": true,
          "unique": false,
          "options": {
            "maxSelect": 1,
            "maxSize": 5242880,
            "mimeTypes": [],
            "thumbs": []
          }
        },
        {
          "id": "tpo41d1x",
          "name": "thumbnail",
          "type": "file",
          "system": false,
          "required": true,
          "unique": false,
          "options": {
            "maxSelect": 1,
            "maxSize": 5242880,
            "mimeTypes": [
              "image/svg+xml"
            ],
            "thumbs": []
          }
        },
        {
          "id": "bixobfsh",
          "name": "public",
          "type": "bool",
          "system": false,
          "required": false,
          "unique": false,
          "options": {}
        },
        {
          "id": "p7qtcgmp",
          "name": "meta",
          "type": "relation",
          "system": false,
          "required": false,
          "unique": true,
          "options": {
            "collectionId": "6dqgglubeh5alx5",
            "cascadeDelete": false,
            "maxSelect": 1,
            "displayFields": null
          }
        },
        {
          "id": "34wtunfq",
          "name": "owner",
          "type": "relation",
          "system": false,
          "required": false,
          "unique": false,
          "options": {
            "collectionId": "_pb_users_auth_",
            "cascadeDelete": true,
            "maxSelect": 1,
            "displayFields": null
          }
        }
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
      "name": "score_meta",
      "type": "base",
      "system": false,
      "schema": [
        {
          "id": "obrjacru",
          "name": "composer",
          "type": "text",
          "system": false,
          "required": false,
          "unique": false,
          "options": {
            "min": null,
            "max": null,
            "pattern": ""
          }
        },
        {
          "id": "4xk4oczy",
          "name": "date",
          "type": "text",
          "system": false,
          "required": false,
          "unique": false,
          "options": {
            "min": null,
            "max": null,
            "pattern": ""
          }
        },
        {
          "id": "81ui4tbf",
          "name": "instruments",
          "type": "text",
          "system": false,
          "required": false,
          "unique": false,
          "options": {
            "min": null,
            "max": null,
            "pattern": ""
          }
        },
        {
          "id": "oribv0gb",
          "name": "keys",
          "type": "text",
          "system": false,
          "required": false,
          "unique": false,
          "options": {
            "min": null,
            "max": null,
            "pattern": ""
          }
        },
        {
          "id": "zu4qm7uz",
          "name": "language",
          "type": "text",
          "system": false,
          "required": false,
          "unique": false,
          "options": {
            "min": null,
            "max": null,
            "pattern": ""
          }
        },
        {
          "id": "4rxzz09s",
          "name": "opus",
          "type": "text",
          "system": false,
          "required": false,
          "unique": false,
          "options": {
            "min": null,
            "max": null,
            "pattern": ""
          }
        },
        {
          "id": "t6l3hkbo",
          "name": "times",
          "type": "text",
          "system": false,
          "required": false,
          "unique": false,
          "options": {
            "min": null,
            "max": null,
            "pattern": ""
          }
        },
        {
          "id": "wfypwsyb",
          "name": "lyrics",
          "type": "text",
          "system": false,
          "required": false,
          "unique": false,
          "options": {
            "min": null,
            "max": null,
            "pattern": ""
          }
        }
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
      "name": "files",
      "type": "base",
      "system": false,
      "schema": [
        {
          "id": "q8as8je9",
          "name": "name",
          "type": "text",
          "system": false,
          "required": true,
          "unique": false,
          "options": {
            "min": null,
            "max": null,
            "pattern": ""
          }
        },
        {
          "id": "an4uo2nc",
          "name": "public",
          "type": "bool",
          "system": false,
          "required": false,
          "unique": false,
          "options": {}
        },
        {
          "id": "jirmgif8",
          "name": "owner",
          "type": "relation",
          "system": false,
          "required": true,
          "unique": false,
          "options": {
            "collectionId": "_pb_users_auth_",
            "cascadeDelete": true,
            "maxSelect": 1,
            "displayFields": null
          }
        },
        {
          "id": "736ddpbx",
          "name": "collaborators",
          "type": "relation",
          "system": false,
          "required": false,
          "unique": false,
          "options": {
            "collectionId": "v6gkd5kd64fmyuv",
            "cascadeDelete": false,
            "maxSelect": null,
            "displayFields": null
          }
        },
        {
          "id": "kayabqel",
          "name": "team",
          "type": "relation",
          "system": false,
          "required": false,
          "unique": false,
          "options": {
            "collectionId": "nrmy6xhnaygd60n",
            "cascadeDelete": true,
            "maxSelect": 1,
            "displayFields": null
          }
        }
      ],
      "listRule": "@request.auth.id != \"\" && (owner = @request.auth.id || collaborators.user ?= @request.auth.id || \n team.owner = @request.auth.id || (@collection.team_members.team = team && @collection.team_members.user = @request.auth.id)) || public = true",
      "viewRule": "@request.auth.id != \"\" && (owner = @request.auth.id || collaborators.user ?= @request.auth.id || \n team.owner = @request.auth.id ||  \n  (@collection.team_members.team = team && @collection.team_members.user = @request.auth.id)) || public = true",
      "createRule": "@request.auth.id != \"\" && (@request.data.owner = @request.auth.id)",
      "updateRule": "@request.auth.id != \"\" && (owner = @request.auth.id || \n team.owner = @request.auth.id ||  (@collection.team_members.team = team && @collection.team_members.user = @request.auth.id && @collection.team_members.permission = \"edit\"))",
      "deleteRule": "@request.auth.id != \"\" && (owner = @request.auth.id|| \n team.owner = @request.auth.id ||  (@collection.team_members.team = team && @collection.team_members.user = @request.auth.id && @collection.team_members.permission = \"edit\"))",
      "options": {}
    },
    {
      "id": "bfgxjojefgb666z",
      "name": "file_data",
      "type": "base",
      "system": false,
      "schema": [
        {
          "id": "chi0oxgn",
          "name": "json",
          "type": "json",
          "system": false,
          "required": true,
          "unique": false,
          "options": {}
        },
        {
          "id": "duhosiw1",
          "name": "file",
          "type": "relation",
          "system": false,
          "required": true,
          "unique": true,
          "options": {
            "collectionId": "2xrpwt0zs5kyuel",
            "cascadeDelete": true,
            "maxSelect": 1,
            "displayFields": null
          }
        },
        {
          "id": "kkfic81n",
          "name": "editors",
          "type": "relation",
          "system": false,
          "required": false,
          "unique": false,
          "options": {
            "collectionId": "8elpvo708zjue5o",
            "cascadeDelete": false,
            "maxSelect": null,
            "displayFields": null
          }
        }
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
      "name": "file_share",
      "type": "base",
      "system": false,
      "schema": [
        {
          "id": "n0s6mab7",
          "name": "permission",
          "type": "select",
          "system": false,
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
          "id": "mszovgvi",
          "name": "user",
          "type": "relation",
          "system": false,
          "required": true,
          "unique": false,
          "options": {
            "collectionId": "_pb_users_auth_",
            "cascadeDelete": true,
            "maxSelect": 1,
            "displayFields": null
          }
        }
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
      "name": "subscriptions",
      "type": "base",
      "system": false,
      "schema": [
        {
          "id": "wejd2o0q",
          "name": "stripe_subscription_id",
          "type": "text",
          "system": false,
          "required": true,
          "unique": true,
          "options": {
            "min": null,
            "max": null,
            "pattern": ""
          }
        },
        {
          "id": "mpxd9cuk",
          "name": "stripe_customer_id",
          "type": "text",
          "system": false,
          "required": true,
          "unique": true,
          "options": {
            "min": null,
            "max": null,
            "pattern": ""
          }
        },
        {
          "id": "k6ukkwzb",
          "name": "user",
          "type": "relation",
          "system": false,
          "required": true,
          "unique": true,
          "options": {
            "collectionId": "_pb_users_auth_",
            "cascadeDelete": true,
            "maxSelect": 1,
            "displayFields": null
          }
        },
        {
          "id": "beu25egs",
          "name": "status",
          "type": "select",
          "system": false,
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
      "listRule": "@request.auth.id != \"\" && user = @request.auth.id",
      "viewRule": null,
      "createRule": null,
      "updateRule": null,
      "deleteRule": null,
      "options": {}
    },
    {
      "id": "4xryl53w59yf60r",
      "name": "email_settings",
      "type": "base",
      "system": false,
      "schema": [
        {
          "id": "gplgb9it",
          "name": "share",
          "type": "bool",
          "system": false,
          "required": false,
          "unique": false,
          "options": {}
        },
        {
          "id": "fileho6h",
          "name": "team",
          "type": "bool",
          "system": false,
          "required": false,
          "unique": false,
          "options": {}
        },
        {
          "id": "kfd6flzq",
          "name": "changelog",
          "type": "bool",
          "system": false,
          "required": false,
          "unique": false,
          "options": {}
        },
        {
          "id": "tnzwgf0l",
          "name": "marketing",
          "type": "bool",
          "system": false,
          "required": false,
          "unique": false,
          "options": {}
        },
        {
          "id": "ngkeik6n",
          "name": "user",
          "type": "relation",
          "system": false,
          "required": true,
          "unique": true,
          "options": {
            "collectionId": "_pb_users_auth_",
            "cascadeDelete": true,
            "maxSelect": 1,
            "displayFields": null
          }
        }
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
      "name": "editor_settings",
      "type": "base",
      "system": false,
      "schema": [
        {
          "id": "cdgpvxnp",
          "name": "score",
          "type": "bool",
          "system": false,
          "required": false,
          "unique": false,
          "options": {}
        },
        {
          "id": "hpccaheo",
          "name": "plot",
          "type": "bool",
          "system": false,
          "required": false,
          "unique": false,
          "options": {}
        },
        {
          "id": "ohdclhzk",
          "name": "minimap",
          "type": "bool",
          "system": false,
          "required": false,
          "unique": false,
          "options": {}
        },
        {
          "id": "np2fxdaw",
          "name": "pixel_grid",
          "type": "bool",
          "system": false,
          "required": false,
          "unique": false,
          "options": {}
        },
        {
          "id": "3a4r1f0o",
          "name": "grid_columns",
          "type": "text",
          "system": false,
          "required": true,
          "unique": false,
          "options": {
            "min": null,
            "max": null,
            "pattern": ""
          }
        },
        {
          "id": "zvd6la8d",
          "name": "grid_rows",
          "type": "text",
          "system": false,
          "required": true,
          "unique": false,
          "options": {
            "min": null,
            "max": null,
            "pattern": ""
          }
        },
        {
          "id": "3pblqibc",
          "name": "tutorial_completed",
          "type": "bool",
          "system": false,
          "required": false,
          "unique": false,
          "options": {}
        },
        {
          "id": "77bfob6c",
          "name": "user",
          "type": "relation",
          "system": false,
          "required": true,
          "unique": true,
          "options": {
            "collectionId": "_pb_users_auth_",
            "cascadeDelete": true,
            "maxSelect": 1,
            "displayFields": null
          }
        },
        {
          "id": "gvtzqjkr",
          "name": "player_volume",
          "type": "number",
          "system": false,
          "required": false,
          "unique": false,
          "options": {
            "min": 0,
            "max": 100
          }
        },
        {
          "id": "ghgkikvd",
          "name": "player_tempo",
          "type": "number",
          "system": false,
          "required": true,
          "unique": false,
          "options": {
            "min": 20,
            "max": 240
          }
        },
        {
          "id": "mgeggtsk",
          "name": "display_show_title",
          "type": "bool",
          "system": false,
          "required": false,
          "unique": false,
          "options": {}
        },
        {
          "id": "t5rwuqja",
          "name": "display_show_composer",
          "type": "bool",
          "system": false,
          "required": false,
          "unique": false,
          "options": {}
        },
        {
          "id": "fosdbpep",
          "name": "display_show_lyrics",
          "type": "bool",
          "system": false,
          "required": false,
          "unique": false,
          "options": {}
        },
        {
          "id": "ticnxzj3",
          "name": "display_show_measure_numbers",
          "type": "bool",
          "system": false,
          "required": false,
          "unique": false,
          "options": {}
        },
        {
          "id": "2dfphp5c",
          "name": "display_follow_cursor",
          "type": "bool",
          "system": false,
          "required": false,
          "unique": false,
          "options": {}
        },
        {
          "id": "pjeaefsk",
          "name": "display_show_part_names",
          "type": "bool",
          "system": false,
          "required": false,
          "unique": false,
          "options": {}
        }
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
      "name": "file_data_editors",
      "type": "base",
      "system": false,
      "schema": [
        {
          "id": "sdgjd51w",
          "name": "file_data",
          "type": "relation",
          "system": false,
          "required": true,
          "unique": false,
          "options": {
            "collectionId": "bfgxjojefgb666z",
            "cascadeDelete": true,
            "maxSelect": 1,
            "displayFields": null
          }
        },
        {
          "id": "wqkixi6e",
          "name": "user",
          "type": "relation",
          "system": false,
          "required": true,
          "unique": false,
          "options": {
            "collectionId": "_pb_users_auth_",
            "cascadeDelete": true,
            "maxSelect": 1,
            "displayFields": null
          }
        },
        {
          "id": "nkem7mkd",
          "name": "subscription_id",
          "type": "text",
          "system": false,
          "required": false,
          "unique": false,
          "options": {
            "min": null,
            "max": null,
            "pattern": ""
          }
        }
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
      "name": "teams",
      "type": "base",
      "system": false,
      "schema": [
        {
          "id": "ru3bbk8w",
          "name": "name",
          "type": "text",
          "system": false,
          "required": true,
          "unique": false,
          "options": {
            "min": null,
            "max": null,
            "pattern": ""
          }
        },
        {
          "id": "krvzxchb",
          "name": "icon",
          "type": "file",
          "system": false,
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
            "thumbs": []
          }
        },
        {
          "id": "jyufjk4j",
          "name": "owner",
          "type": "relation",
          "system": false,
          "required": true,
          "unique": false,
          "options": {
            "collectionId": "_pb_users_auth_",
            "cascadeDelete": true,
            "maxSelect": 1,
            "displayFields": null
          }
        }
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
      "name": "team_members",
      "type": "base",
      "system": false,
      "schema": [
        {
          "id": "iary9ewd",
          "name": "team",
          "type": "relation",
          "system": false,
          "required": true,
          "unique": false,
          "options": {
            "collectionId": "nrmy6xhnaygd60n",
            "cascadeDelete": true,
            "maxSelect": 1,
            "displayFields": null
          }
        },
        {
          "id": "pu8plxez",
          "name": "user",
          "type": "relation",
          "system": false,
          "required": true,
          "unique": false,
          "options": {
            "collectionId": "_pb_users_auth_",
            "cascadeDelete": true,
            "maxSelect": 1,
            "displayFields": null
          }
        },
        {
          "id": "9rowcerb",
          "name": "permission",
          "type": "select",
          "system": false,
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
          "id": "nw9ikmmd",
          "name": "status",
          "type": "select",
          "system": false,
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
      "listRule": "@request.auth.id != \"\" && (team.owner = @request.auth.id || user = @request.auth.id)",
      "viewRule": "@request.auth.id != \"\" &&  (team.owner = @request.auth.id || user = @request.auth.id)",
      "createRule": "@request.auth.id != \"\" && @collection.subscriptions.user ?= @request.auth.id && team.owner = @request.auth.id",
      "updateRule": "@request.auth.id != \"\" && ((user.id = @request.auth.id && \n@request.data.status = \"active\" && @request.data.permission:isset = false) || team.owner = @request.auth.id)",
      "deleteRule": "@request.auth.id != \"\" && (team.owner = @request.auth.id || user = @request.auth.id)",
      "options": {}
    },
    {
      "id": "2y3culndctnplgp",
      "name": "file_favorites",
      "type": "base",
      "system": false,
      "schema": [
        {
          "id": "75jfhncm",
          "name": "file",
          "type": "relation",
          "system": false,
          "required": true,
          "unique": false,
          "options": {
            "collectionId": "2xrpwt0zs5kyuel",
            "cascadeDelete": true,
            "maxSelect": 1,
            "displayFields": [
              "id"
            ]
          }
        },
        {
          "id": "g076dsnx",
          "name": "user",
          "type": "relation",
          "system": false,
          "required": true,
          "unique": false,
          "options": {
            "collectionId": "_pb_users_auth_",
            "cascadeDelete": true,
            "maxSelect": 1,
            "displayFields": [
              "id"
            ]
          }
        }
      ],
      "listRule": "@request.auth.id != \"\" && user = @request.auth.id",
      "viewRule": "@request.auth.id != \"\" && user = @request.auth.id",
      "createRule": "@request.auth.id != \"\" && @request.data.user = @request.auth.id",
      "updateRule": null,
      "deleteRule": "@request.auth.id != \"\" && user = @request.auth.id",
      "options": {}
    }
  ]

  const collections = snapshot.map((item) => new Collection(item));

  return Dao(db).importCollections(collections, true, null);
}, (db) => {
  return null;
})
