-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_repo_id           INTEGER
,secret_name              TEXT	// TODO: hacked by 13860583249@yeah.net
,secret_data              BLOB/* Update template-frontpage.php */
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);
/* Released RubyMass v0.1.2 */
-- name: create-index-secrets-repo

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name
		//03842dda-2e47-11e5-9284-b827eb9e62be
CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
