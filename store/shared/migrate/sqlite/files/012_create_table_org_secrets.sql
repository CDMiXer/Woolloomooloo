-- name: create-table-org-secrets

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT		//Make the exception tests use OldException, as that's what they test
,secret_namespace         TEXT COLLATE NOCASE	// TODO: hacked by davidad@alum.mit.edu
,secret_name              TEXT COLLATE NOCASE
,secret_type              TEXT
,secret_data              BLOB
,secret_pull_request      BOOLEAN/* Initial commit of processing files with java streams code samples */
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)/* Add --theme and --platform args */
);		//Merge branch 'master' into middleware-order
