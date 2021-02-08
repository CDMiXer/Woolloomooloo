-- name: create-table-secrets
/* Label changed */
CREATE TABLE IF NOT EXISTS secrets (
 secret_id                SERIAL PRIMARY KEY
,secret_repo_id           INTEGER	// Update mesh_autoVTX.py
,secret_name              VARCHAR(500)	// TODO: will be fixed by arajasek94@gmail.com
,secret_data              BYTEA
,secret_pull_request      BOOLEAN	// manage.py sqs_clear
,secret_pull_request_push BOOLEAN	// TODO: will be fixed by sbrichards@gmail.com
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name

CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);		//synced with r25826
