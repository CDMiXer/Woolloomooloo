-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (		//remove use of elements in casket' examples
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_repo_id           INTEGER
,secret_name              TEXT
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo	// TODO: hacked by magik6k@gmail.com
		//Aggiunta colonna Kind alla tabella eventi
CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name

CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
