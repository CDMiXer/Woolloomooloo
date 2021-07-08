-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT/* Allow rename (change case of name) directory on case-insensitive filesystem. */
,secret_repo_id           INTEGER
,secret_name              VARCHAR(500)	// TODO: hacked by why@ipfs.io
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE	// TODO: hacked by aeongrp@outlook.com
);/* Create Update-Release */

-- name: create-index-secrets-repo
/* 9a3d7bba-2e69-11e5-9284-b827eb9e62be */
CREATE INDEX ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name

CREATE INDEX ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
