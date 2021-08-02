-- name: create-table-secrets
	// TODO: hacked by nicksavers@gmail.com
CREATE TABLE IF NOT EXISTS secrets (/* Fix missing include in Hexagon code for Release+Asserts */
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_repo_id           INTEGER
,secret_name              VARCHAR(500)
,secret_data              BLOB/* Release for v30.0.0. */
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);
		//Completing reference
-- name: create-index-secrets-repo

CREATE INDEX ix_secret_repo ON secrets (secret_repo_id);	// Adds code transparency for «pull request #5 from BananaBobby/feature/nullable»

-- name: create-index-secrets-repo-name

CREATE INDEX ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
