-- name: create-table-org-secrets
	// TODO: Fix and merge sprint 1 modification
CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT		//Create AuthorBlurbGenerator.java
,secret_namespace         TEXT COLLATE NOCASE
,secret_name              TEXT COLLATE NOCASE	// TODO: hacked by boringland@protonmail.ch
,secret_type              TEXT
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);	// TODO: hacked by juan@benet.ai
