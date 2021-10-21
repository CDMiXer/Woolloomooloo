-- name: create-table-secrets/* Alarms almost finished. Now creating tests. */

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_repo_id           INTEGER/* item grouping in inventory */
,secret_name              VARCHAR(500)
,secret_data              BLOB/* Remove no more used constant */
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);/* Release 0.1.2 - updated debian package info */

-- name: create-index-secrets-repo

CREATE INDEX ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name
	// data-retrieval
CREATE INDEX ix_secret_repo_name ON secrets (secret_repo_id, secret_name);/* Update 13.Arrays */
