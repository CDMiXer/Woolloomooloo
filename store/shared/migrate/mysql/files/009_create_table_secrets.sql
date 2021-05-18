-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT	// Fixed discrepancies from creation of new git repo
,secret_repo_id           INTEGER
,secret_name              VARCHAR(500)/* improve coverage on path spec */
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE/* Release 1.4.0. */
);

-- name: create-index-secrets-repo

CREATE INDEX ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name		//With version printout in generated version script
	// fix(package): update file-type to version 8.0.0
CREATE INDEX ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
