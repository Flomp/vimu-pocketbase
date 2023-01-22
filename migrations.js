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
            "maxSelect": 1,
            "collectionId": "6dqgglubeh5alx5",
            "cascadeDelete": false
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
            "maxSelect": 1,
            "collectionId": "_pb_users_auth_",
            "cascadeDelete": true
          }
        }
      ],
      "listRule": "@request.auth.id != \"\" && (owner = @request.auth.id || public = true)",
      "viewRule": "@request.auth.id != \"\" && (owner = @request.auth.id || public = true)",
      "createRule": "",
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
      "viewRule": "@request.auth.id != \"\" && @collection.scores.meta = id && (@collection.scores.owner = @request.auth.id || @collection.scores.public = true)",
      "createRule": "",
      "updateRule": "@request.auth.id != \"\" && @collection.scores.meta = id && (@collection.scores.owner = @request.auth.id)",
      "deleteRule": "@request.auth.id != \"\" && @collection.scores.meta = id && (@collection.scores.owner = @request.auth.id)",
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
          "id": "imchaigi",
          "name": "favorite",
          "type": "bool",
          "system": false,
          "required": false,
          "unique": false,
          "options": {}
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
            "maxSelect": 1,
            "collectionId": "_pb_users_auth_",
            "cascadeDelete": true
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
            "maxSelect": null,
            "collectionId": "v6gkd5kd64fmyuv",
            "cascadeDelete": false
          }
        }
      ],
      "listRule": "@request.auth.id != \"\" && (owner = @request.auth.id || collaborators.user ?~ @request.auth.id)",
      "viewRule": "@request.auth.id != \"\" && (owner = @request.auth.id || collaborators.user ?~ @request.auth.id || public = true)",
      "createRule": "@request.auth.id != \"\" && (owner = @request.auth.id)",
      "updateRule": "@request.auth.id != \"\" && (owner = @request.auth.id)",
      "deleteRule": "@request.auth.id != \"\" && (owner = @request.auth.id)",
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
            "maxSelect": 1,
            "collectionId": "2xrpwt0zs5kyuel",
            "cascadeDelete": true
          }
        }
      ],
      "listRule": "@request.auth.id != \"\" && file.owner = @request.auth.id",
      "viewRule": "@request.auth.id != \"\" && (file.owner = @request.auth.id || file.collaborators.user.id ?= @request.auth.id)",
      "createRule": "@request.auth.id != \"\" && file.owner = @request.auth.id",
      "updateRule": "@request.auth.id != \"\" && (file.owner = @request.auth.id || (file.collaborators.user.id ?= @request.auth.id && file.collaborators.permission = \"edit\"))",
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
            "maxSelect": 1,
            "collectionId": "_pb_users_auth_",
            "cascadeDelete": true
          }
        }
      ],
      "listRule": "@request.auth.id != \"\" && (@collection.files.collaborators.id ?= id && (@collection.files.owner = @request.auth.id || user = @request.auth.id))",
      "viewRule": "@request.auth.id != \"\" && (@collection.files.collaborators.id ?= id && (@collection.files.owner = @request.auth.id || user = @request.auth.id))",
      "createRule": "@request.auth.id != \"\" && @collection.subscriptions.user = @request.auth.id",
      "updateRule": "@request.auth.id != \"\" && (@collection.files.collaborators.id ?= id && @collection.files.owner = @request.auth.id) && @collection.subscriptions.user = @request.auth.id",
      "deleteRule": "@request.auth.id != \"\" && (@collection.files.collaborators.id ?= id && @collection.files.owner = @request.auth.id) && @collection.subscriptions.user = @request.auth.id",
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
            "maxSelect": 1,
            "collectionId": "_pb_users_auth_",
            "cascadeDelete": true
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
    }
  ];

  const collections = snapshot.map((item) => new Collection(item));

  return Dao(db).importCollections(collections, true, null);
}, (db) => {
  return null;
})
