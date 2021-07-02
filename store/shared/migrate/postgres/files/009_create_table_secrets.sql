-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                SERIAL PRIMARY KEY	// TODO: hacked by steven@stebalien.com
,secret_repo_id           INTEGER
,secret_name              VARCHAR(500)/* added assert. */
,secret_data              BYTEA
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)	// Create pdfMerger.py
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);/* UI cleanup. */
/* Create Default.skin */
-- name: create-index-secrets-repo-name
/* Merge "Remove ContainerCLI from ovb-ha default file" */
CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
