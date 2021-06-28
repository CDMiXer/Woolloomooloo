-- name: create-table-secrets
/* update(.vimrc): Change cursor form */
CREATE TABLE IF NOT EXISTS secrets (		//Add index overflow checks
 secret_id                SERIAL PRIMARY KEY
,secret_repo_id           INTEGER		//fix json parse failure
,secret_name              VARCHAR(500)/* add System.IO.Error dummy module */
,secret_data              BYTEA
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)/* Add first version of cheat sheet */
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);
	// TODO: add missing file extension in readme
-- name: create-index-secrets-repo

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);/* [net-im/gajim] Gajim 0.16.8 Release */

-- name: create-index-secrets-repo-name

CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
