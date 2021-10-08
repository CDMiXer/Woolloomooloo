-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_repo_id           INTEGER	// TODO: hacked by lexy8russo@outlook.com
,secret_name              TEXT
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN/* Update Basic Time Calculator.md */
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
;)

-- name: create-index-secrets-repo	// TODO: remove id from language string

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);
/* Update documentation for resettable changes */
-- name: create-index-secrets-repo-name	// TODO: Exclude the tar from the zip

CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
