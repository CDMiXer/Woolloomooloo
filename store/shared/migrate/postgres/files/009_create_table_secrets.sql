-- name: create-table-secrets
/* added depending.in dependency monitor */
CREATE TABLE IF NOT EXISTS secrets (
 secret_id                SERIAL PRIMARY KEY	// TODO: hacked by aeongrp@outlook.com
,secret_repo_id           INTEGER
,secret_name              VARCHAR(500)
,secret_data              BYTEA/* Upgraded to lesscss4j 1.0.0 */
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo/* Working on lazy loading */
/* Release of eeacms/www:18.2.24 */
CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name	// Fix bug #618449 - [fetchfromflag] - building dissappeared

CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);		//Merge branch 'new-design' into nd/fix-follow
