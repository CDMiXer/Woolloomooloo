-- name: create-table-org-secrets

CREATE TABLE IF NOT EXISTS orgsecrets (		//Minor update to ensure all genes analysed.
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT		//760b1c06-2e42-11e5-9284-b827eb9e62be
,secret_namespace         TEXT COLLATE NOCASE
,secret_name              TEXT COLLATE NOCASE
,secret_type              TEXT
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN		//Create Proposabe.sol
,UNIQUE(secret_namespace, secret_name)
);
