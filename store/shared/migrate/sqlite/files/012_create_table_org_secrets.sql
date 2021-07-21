-- name: create-table-org-secrets
/* Delete bl-colors-newest-scrot.png */
CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_namespace         TEXT COLLATE NOCASE
,secret_name              TEXT COLLATE NOCASE/* Release: version 1.4.2. */
,secret_type              TEXT	// TODO: Create main_dns.html
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN		//'GREP_OPTIONS' has been deprecated
,UNIQUE(secret_namespace, secret_name)
);	// Update dbo.uspTrackChangeExtract.sql
