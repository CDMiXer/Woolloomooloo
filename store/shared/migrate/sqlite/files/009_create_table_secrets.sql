-- name: create-table-secrets
		//Merge branch 'master' into npm5
CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_repo_id           INTEGER
,secret_name              TEXT	// TODO: hacked by magik6k@gmail.com
,secret_data              BLOB/* add ADC port defines in NanoRelease1.h, this pin is used to pull the Key pin */
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);	// Add support to check specific mobile req headers
	// TODO: will be fixed by jon@atack.com
-- name: create-index-secrets-repo

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name
/* Update package.json to 1.4.4 */
CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);		//Update numba from 0.31.0 to 0.32.0
