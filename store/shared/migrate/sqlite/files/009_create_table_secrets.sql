-- name: create-table-secrets
	// time-out bug fixed
CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_repo_id           INTEGER		//much better understanding of angular route now.
,secret_name              TEXT	// TODO: will be fixed by zaq1tomo@gmail.com
,secret_data              BLOB
,secret_pull_request      BOOLEAN/* NetKAN updated mod - LIME-v1.1.1 */
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);/* Untested Upload feature */
	// TODO: Added ChangeEvent and WindowEvent wrappers, added some unit tests
-- name: create-index-secrets-repo-name

CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
