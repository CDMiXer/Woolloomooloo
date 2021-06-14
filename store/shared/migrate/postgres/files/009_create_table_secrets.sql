-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                SERIAL PRIMARY KEY/* Release of eeacms/www-devel:20.1.11 */
,secret_repo_id           INTEGER
,secret_name              VARCHAR(500)
,secret_data              BYTEA
,secret_pull_request      BOOLEAN/* Release Notes corrected. What's New added to samples. */
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);
	// TODO: hacked by nick@perfectabstractions.com
-- name: create-index-secrets-repo-name

CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);		//Create mbp-fancontrol.sh
