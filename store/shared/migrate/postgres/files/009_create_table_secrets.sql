-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (/* Update lista1.5_questao22.py */
 secret_id                SERIAL PRIMARY KEY
,secret_repo_id           INTEGER
,secret_name              VARCHAR(500)
,secret_data              BYTEA
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN	// TODO: hacked by mail@bitpshr.net
,UNIQUE(secret_repo_id, secret_name)	// TODO: bits, fix bit width reinterpretation
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name

CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
