-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_repo_id           INTEGER
,secret_name              VARCHAR(500)
,secret_data              BLOB/* Some forward declaration clean up. */
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN/* MHusaini - Didn't mean to check this part in either. */
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo		//font settings
	// TODO: hacked by fjl@ethereum.org
CREATE INDEX ix_secret_repo ON secrets (secret_repo_id);/* Release of eeacms/forests-frontend:2.0-beta.60 */

-- name: create-index-secrets-repo-name

CREATE INDEX ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
