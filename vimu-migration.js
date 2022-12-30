migrate((db) => {
  const snapshot = [
    {
      "id": "_pb_users_auth_",
      "created": "2022-12-29 16:10:40.503Z",
      "updated": "2022-12-29 16:11:47.978Z",
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
        "requireEmail": false
      }
    },
    {
      "id": "2xrpwt0zs5kyuel",
      "created": "2022-12-29 16:11:47.979Z",
      "updated": "2022-12-29 16:11:47.979Z",
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
          "id": "imchaigi",
          "name": "favorite",
          "type": "bool",
          "required": false,
          "unique": false,
          "options": {}
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
          "id": "qd86rdyl",
          "name": "data",
          "type": "relation",
          "required": false,
          "unique": false,
          "options": {
            "maxSelect": 1,
            "collectionId": "bfgxjojefgb666z",
            "cascadeDelete": false
          }
        },
        {
          "system": false,
          "id": "jirmgif8",
          "name": "owner",
          "type": "relation",
          "required": true,
          "unique": false,
          "options": {
            "maxSelect": 1,
            "collectionId": "_pb_users_auth_",
            "cascadeDelete": true
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
            "maxSelect": null,
            "collectionId": "v6gkd5kd64fmyuv",
            "cascadeDelete": false
          }
        }
      ],
      "listRule": "@request.auth.id != \"\" && (owner = @request.auth.id || collaborators.user ~ @request.auth.id)",
      "viewRule": "@request.auth.id != \"\" && (owner = @request.auth.id || collaborators.user ~ @request.auth.id || public = true)",
      "createRule": "@request.auth.id != \"\"",
      "updateRule": "@request.auth.id != \"\" && (owner = @request.auth.id)",
      "deleteRule": "@request.auth.id != \"\" && (owner = @request.auth.id)",
      "options": {}
    },
    {
      "id": "bfgxjojefgb666z",
      "created": "2022-12-29 16:11:47.979Z",
      "updated": "2022-12-29 16:11:47.979Z",
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
        }
      ],
      "listRule": "@request.auth.id != \"\" && @collection.files.data = id && @collection.files.owner = @request.auth.id",
      "viewRule": "@request.auth.id != \"\" && @collection.files.data = id && (@collection.files.owner = @request.auth.id || @collection.files.collaborators.user.id = @request.auth.id)",
      "createRule": "",
      "updateRule": "@request.auth.id != \"\" && @collection.files.data = id && (@collection.files.owner = @request.auth.id || (@collection.files.collaborators.user.id = @request.auth.id && @collection.files.collaborators.permission = \"edit\"))",
      "deleteRule": "@request.auth.id != \"\" && @collection.files.data = id && @collection.files.owner = @request.auth.id",
      "options": {}
    },
    {
      "id": "v6gkd5kd64fmyuv",
      "created": "2022-12-29 16:11:47.979Z",
      "updated": "2022-12-29 16:11:47.979Z",
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
            "maxSelect": 1,
            "collectionId": "_pb_users_auth_",
            "cascadeDelete": true
          }
        }
      ],
      "listRule": "@request.auth.id != \"\" && (@collection.files.collaborators.id = id && (@collection.files.owner = @request.auth.id || user = @request.auth.id))",
      "viewRule": "@request.auth.id != \"\" && (@collection.files.collaborators.id = id && (@collection.files.owner = @request.auth.id || user = @request.auth.id))",
      "createRule": "",
      "updateRule": "@request.auth.id != \"\" && (@collection.files.collaborators.id = id && @collection.files.owner = @request.auth.id)",
      "deleteRule": "@request.auth.id != \"\" && (@collection.files.collaborators.id = id && @collection.files.owner = @request.auth.id)",
      "options": {}
    },
    {
      "id": "1b8m8nuntgbyzal",
      "created": "2022-12-29 16:11:47.980Z",
      "updated": "2022-12-29 16:11:47.980Z",
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
            "thumbs": []
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
            "thumbs": []
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
          "id": "34wtunfq",
          "name": "owner",
          "type": "relation",
          "required": false,
          "unique": false,
          "options": {
            "maxSelect": 1,
            "collectionId": "_pb_users_auth_",
            "cascadeDelete": true
          }
        },
        {
          "system": false,
          "id": "tay3uruk",
          "name": "meta",
          "type": "relation",
          "required": false,
          "unique": false,
          "options": {
            "maxSelect": 1,
            "collectionId": "6dqgglubeh5alx5",
            "cascadeDelete": false
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
      "created": "2022-12-29 16:11:47.980Z",
      "updated": "2022-12-29 16:11:47.980Z",
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
        }
      ],
      "listRule": "@request.auth.id != \"\" && @collection.scores.meta = id && (@collection.scores.owner = @request.auth.id || @collection.scores.public = true)",
      "viewRule": "@request.auth.id != \"\" && @collection.scores.meta = id && (@collection.scores.owner = @request.auth.id || @collection.scores.public = true)",
      "createRule": "",
      "updateRule": "@request.auth.id != \"\" && @collection.scores.meta = id && (@collection.scores.owner = @request.auth.id)",
      "deleteRule": "@request.auth.id != \"\" && @collection.scores.meta = id && (@collection.scores.owner = @request.auth.id)",
      "options": {}
    }
  ];

  const collections = snapshot.map((item) => new Collection(item));

  return Dao(db).importCollections(collections, true, null);
}, (db) => {
  return null;
})
