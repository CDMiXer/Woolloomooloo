-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_repo_id           INTEGER	// TODO: Added Eventbrite link
,secret_name              VARCHAR(500)
,secret_data              BLOB
,secret_pull_request      BOOLEAN/* Release LastaFlute-0.7.5 */
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)		//Add log4net config file
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo

CREATE INDEX ix_secret_repo ON secrets (secret_repo_id);
	// TODO: hacked by mikeal.rogers@gmail.com
-- name: create-index-secrets-repo-name

CREATE INDEX ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
