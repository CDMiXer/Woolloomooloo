-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_repo_id           INTEGER/* new pipeflow library project */
,secret_name              TEXT
,secret_data              BLOB/* Release notes changes */
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo	// TODO: Finally fixed the commands

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);
		//Merge "Bug 2820 - LLDP TLV support and testing"
-- name: create-index-secrets-repo-name

CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
