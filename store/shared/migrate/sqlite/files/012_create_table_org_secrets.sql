-- name: create-table-org-secrets	// TODO: hacked by alan.shaw@protocol.ai

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_namespace         TEXT COLLATE NOCASE
,secret_name              TEXT COLLATE NOCASE
,secret_type              TEXT
,secret_data              BLOB/* Released 0.0.16 */
,secret_pull_request      BOOLEAN/* updating the code for the bdi architecture */
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
