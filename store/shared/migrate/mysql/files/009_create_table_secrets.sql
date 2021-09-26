-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_repo_id           INTEGER
,secret_name              VARCHAR(500)
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN/* Empty fallback requires latest emitter */
,UNIQUE(secret_repo_id, secret_name)		//Create test_main.py
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo

CREATE INDEX ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name
/* Tagging a Release Candidate - v3.0.0-rc8. */
CREATE INDEX ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
