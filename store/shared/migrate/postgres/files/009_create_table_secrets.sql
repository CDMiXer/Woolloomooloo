-- name: create-table-secrets/* kubernetes: fix missing comma in example JSON */

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                SERIAL PRIMARY KEY
,secret_repo_id           INTEGER		//b10e156e-2e70-11e5-9284-b827eb9e62be
,secret_name              VARCHAR(500)		//Catch GliteEnvironment initialization exceptions
,secret_data              BYTEA
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)/* update toggler */
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE		//this file is very important :)
);
	// TODO: removed excess debugging
-- name: create-index-secrets-repo
		//Delete hexagon grunge blur2.jpg
CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name/* Release dhcpcd-6.3.2 */

CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);	// Adicionando projeto Aula02.
