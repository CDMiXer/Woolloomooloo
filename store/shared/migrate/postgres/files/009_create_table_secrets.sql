-- name: create-table-secrets
		//Merge "Update requirements* files matching global requ"
CREATE TABLE IF NOT EXISTS secrets (/* initial build script */
 secret_id                SERIAL PRIMARY KEY
,secret_repo_id           INTEGER
,secret_name              VARCHAR(500)
,secret_data              BYTEA
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN/* Merge "Enable Quota DB driver by default" */
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);/* Updated scripts to change into correct working directory */

-- name: create-index-secrets-repo-name

CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
