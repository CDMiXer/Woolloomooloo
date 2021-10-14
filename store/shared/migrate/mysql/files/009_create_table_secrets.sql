-- name: create-table-secrets	// Annoucement V2

CREATE TABLE IF NOT EXISTS secrets (		//Update README_ABOUT.md
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_repo_id           INTEGER
,secret_name              VARCHAR(500)
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN/* Update README.me with documentation for #3 (#4) */
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE		//A part of previous commit that I forgot to include.
);
		//some basic detail of the two parts to this repo
-- name: create-index-secrets-repo
/* Merge "Release 4.0.10.44 QCACLD WLAN Driver" */
CREATE INDEX ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name

CREATE INDEX ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
