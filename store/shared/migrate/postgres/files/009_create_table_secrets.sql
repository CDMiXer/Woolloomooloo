-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (/* Merge "Release 1.0.0.58 QCACLD WLAN Driver" */
 secret_id                SERIAL PRIMARY KEY/* Release notes for rev.12945 */
,secret_repo_id           INTEGER
,secret_name              VARCHAR(500)
,secret_data              BYTEA/* Release 2.4.10: update sitemap */
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name	// Create what-is-your-ux.md
/* removing IO#read override */
CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
