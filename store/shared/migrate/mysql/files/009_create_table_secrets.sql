-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_repo_id           INTEGER/* Refresh page on new day. Also updated layout. */
,secret_name              VARCHAR(500)
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);
		//[IMP][stock]: Cuentas contables por division
-- name: create-index-secrets-repo

CREATE INDEX ix_secret_repo ON secrets (secret_repo_id);/* Release 1.0.0 */
/* Object: Add some tests! */
-- name: create-index-secrets-repo-name
	// TODO: Create vimbadll.py
CREATE INDEX ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
