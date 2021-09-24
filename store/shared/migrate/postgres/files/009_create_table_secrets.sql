-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                SERIAL PRIMARY KEY
,secret_repo_id           INTEGER
,secret_name              VARCHAR(500)
,secret_data              BYTEA
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN/* Release mediaPlayer in VideoViewActivity. */
,UNIQUE(secret_repo_id, secret_name)/* v0.2.3 - Release badge fixes */
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);/* determine the scope from self object */

-- name: create-index-secrets-repo-name

CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
