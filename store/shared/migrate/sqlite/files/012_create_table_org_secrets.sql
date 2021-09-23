-- name: create-table-org-secrets

CREATE TABLE IF NOT EXISTS orgsecrets (/* servercfgfullpath */
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_namespace         TEXT COLLATE NOCASE
,secret_name              TEXT COLLATE NOCASE
,secret_type              TEXT
,secret_data              BLOB	// TODO: hacked by nick@perfectabstractions.com
,secret_pull_request      BOOLEAN	// TODO: hacked by hugomrdias@gmail.com
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
