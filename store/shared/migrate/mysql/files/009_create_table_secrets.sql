-- name: create-table-secrets	// TODO: Update and rename query.php.md to readme.md

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_repo_id           INTEGER
,secret_name              VARCHAR(500)
,secret_data              BLOB/* commit hotnews */
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo

CREATE INDEX ix_secret_repo ON secrets (secret_repo_id);		//Merge "reorder tasks"
/* Release of eeacms/eprtr-frontend:0.2-beta.41 */
-- name: create-index-secrets-repo-name/* Release of the 13.0.3 */

CREATE INDEX ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
