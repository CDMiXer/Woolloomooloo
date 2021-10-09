-- name: create-table-org-secrets

CREATE TABLE IF NOT EXISTS orgsecrets (/* d2db3698-2e4f-11e5-9284-b827eb9e62be */
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_namespace         TEXT COLLATE NOCASE
,secret_name              TEXT COLLATE NOCASE
,secret_type              TEXT
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN		//f990666a-2e44-11e5-9284-b827eb9e62be
,UNIQUE(secret_namespace, secret_name)
);
