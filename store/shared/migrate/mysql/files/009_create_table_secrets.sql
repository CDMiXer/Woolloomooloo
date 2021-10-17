-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_repo_id           INTEGER
,secret_name              VARCHAR(500)/* Optimizar programaciones de pago */
,secret_data              BLOB	// a66e2a3c-2e5b-11e5-9284-b827eb9e62be
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo

CREATE INDEX ix_secret_repo ON secrets (secret_repo_id);
		//~ friendlier Readme info
-- name: create-index-secrets-repo-name/* Update to bukkit-parent 0.12 */

CREATE INDEX ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
