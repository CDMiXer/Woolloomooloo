-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                SERIAL PRIMARY KEY
,secret_repo_id           INTEGER	// TODO: Added Usability Report
,secret_name              VARCHAR(500)
,secret_data              BYTEA		//Merge pull request #109 from fkautz/pr_out_minor_code_cleanup
,secret_pull_request      BOOLEAN/* TYPO3 CMS 6 Release (v1.0.0) */
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);		//1.0.99-RC1

-- name: create-index-secrets-repo

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name

CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
