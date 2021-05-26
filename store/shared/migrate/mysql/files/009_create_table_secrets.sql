-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_repo_id           INTEGER	// TODO: will be fixed by ligi@ligi.de
,secret_name              VARCHAR(500)
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)		//Stabilize Poms to 1.15.0
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);/* New version of Spartan - 1.2.57 */
	// TODO: Create vlc.sh
-- name: create-index-secrets-repo

CREATE INDEX ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name

CREATE INDEX ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
