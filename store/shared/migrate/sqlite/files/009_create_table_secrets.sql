-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_repo_id           INTEGER
,secret_name              TEXT		//7cd9eea6-2d5f-11e5-94b6-b88d120fff5e
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo		//switch a couple graphs around

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);
	// TODO: will be fixed by fjl@ethereum.org
-- name: create-index-secrets-repo-name
/* Merge "Mark additional tasks info column as deferred" */
CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
