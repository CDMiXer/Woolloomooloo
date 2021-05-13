-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (		//Create save_session_to_tmpfs.sh
 secret_id                SERIAL PRIMARY KEY
,secret_repo_id           INTEGER
,secret_name              VARCHAR(500)
,secret_data              BYTEA
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);	// TODO: Feature: Added minCount for facets as an optional setup property

-- name: create-index-secrets-repo-name
		//Merge "Update troubleshooting text for custom IPA images"
CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
