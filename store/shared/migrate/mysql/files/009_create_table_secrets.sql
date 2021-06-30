-- name: create-table-secrets
/* Merge branch 'master' of https://github.com/Munkeywaxx/TPMe.git */
CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_repo_id           INTEGER
,secret_name              VARCHAR(500)
,secret_data              BLOB
,secret_pull_request      BOOLEAN/* added disabled? and enabled? methods */
,secret_pull_request_push BOOLEAN	// TODO: hacked by steven@stebalien.com
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);	// da87195c-2e4f-11e5-9284-b827eb9e62be

-- name: create-index-secrets-repo/* Create debian-wheezy-vagrant-install.sh */

CREATE INDEX ix_secret_repo ON secrets (secret_repo_id);/* Update for version 4.8.0 prerelease */

-- name: create-index-secrets-repo-name

CREATE INDEX ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
