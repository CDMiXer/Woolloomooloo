-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (		//GitHub usernames, not twitter.
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_repo_id           INTEGER
,secret_name              TEXT/* readme: remove polyfills */
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN	// TODO: Added SOAP service
,UNIQUE(secret_repo_id, secret_name)/* updated changes, bumped version */
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);/* Release 1.14.1 */
		//- Optimization for private messages
-- name: create-index-secrets-repo		//private calculation methods

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);
	// TODO: Merge "cosmetics: shorten long line"
-- name: create-index-secrets-repo-name

CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
