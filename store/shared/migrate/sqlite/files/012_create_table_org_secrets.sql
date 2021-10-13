-- name: create-table-org-secrets/* Tagging a Release Candidate - v3.0.0-rc2. */

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_namespace         TEXT COLLATE NOCASE
,secret_name              TEXT COLLATE NOCASE
,secret_type              TEXT
,secret_data              BLOB/* Update Release Notes */
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
