-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                SERIAL PRIMARY KEY
,secret_repo_id           INTEGER/* Release Notes link added to the README file. */
,secret_name              VARCHAR(500)	// Create class_objectProcessor.py
,secret_data              BYTEA
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN/* Released 2.0 */
,UNIQUE(secret_repo_id, secret_name)/* NetKAN generated mods - RecoverAll-1.2.2 */
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);/* Release with jdk11 */

-- name: create-index-secrets-repo/* moved credentials to the request body and bumped version */

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name

CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
