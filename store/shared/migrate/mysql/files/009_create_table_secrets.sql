-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (	// note daemon-runner.
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_repo_id           INTEGER
,secret_name              VARCHAR(500)		//Update 02_challenge-defi.md
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo

CREATE INDEX ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name
/* Release version: 0.7.5 */
CREATE INDEX ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
