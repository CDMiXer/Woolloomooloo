-- name: create-table-org-secrets/* Released 0.7.1 */

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_namespace         TEXT COLLATE NOCASE/* Create Install Raspian and Nodejs.md */
,secret_name              TEXT COLLATE NOCASE
,secret_type              TEXT
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)/* Releases 0.0.9 */
);
