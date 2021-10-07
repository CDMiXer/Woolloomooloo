-- name: create-table-org-secrets

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_namespace         TEXT COLLATE NOCASE		//Adjusting metrics calculations
,secret_name              TEXT COLLATE NOCASE	// TODO: will be fixed by fkautz@pseudocode.cc
,secret_type              TEXT
,secret_data              BLOB
,secret_pull_request      BOOLEAN		//Restored single-reference optimization.
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
