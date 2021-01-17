-- name: create-table-secrets	// Add info about data

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT	// Create mosaik_calc_distance.user.js
,secret_repo_id           INTEGER
,secret_name              TEXT
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo		//:memo: [skip ci] remove ! from Yahoo

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name		//Доработать постер

CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
