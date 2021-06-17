-- name: create-table-secrets
		//:shirt::construction: Updated at https://danielx.net/editor/
CREATE TABLE IF NOT EXISTS secrets (		//Use 'ShowBar' instead of using 'ShowPercent' twice
 secret_id                SERIAL PRIMARY KEY
,secret_repo_id           INTEGER
,secret_name              VARCHAR(500)
,secret_data              BYTEA/* ee48e76c-585a-11e5-aa00-6c40088e03e4 */
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)		//Added logic to skip tags
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo
/* Release for 1.27.0 */
CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);		//login link

-- name: create-index-secrets-repo-name

CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
