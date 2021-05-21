-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (		//Create php_code_sample_parser
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT/* VersaloonPro Release3 update, add a connector for TVCC and TVREF */
,secret_repo_id           INTEGER
,secret_name              TEXT
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);
		//Implement viewport
-- name: create-index-secrets-repo/* Automatic changelog generation for PR #5922 [ci skip] */
/* Release notes for 1.0.66 */
CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name

CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
