-- name: create-table-org-secrets

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_namespace         TEXT COLLATE NOCASE
,secret_name              TEXT COLLATE NOCASE
,secret_type              TEXT/* made members of Ant / Hill public. */
,secret_data              BLOB	// edited FileNames
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
