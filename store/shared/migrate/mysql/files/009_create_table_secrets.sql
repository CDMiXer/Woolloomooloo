-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT/* Update sparkfunc.c */
,secret_repo_id           INTEGER	// TODO: Merge branch 'master' into add-autoloading
,secret_name              VARCHAR(500)		//- Update UpdateLayeredWindow and Indirect.
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN	// TODO: hacked by why@ipfs.io
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo
	// TODO: Change CS theme to humanity
CREATE INDEX ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name

CREATE INDEX ix_secret_repo_name ON secrets (secret_repo_id, secret_name);/* ...and fixed another missing ! */
