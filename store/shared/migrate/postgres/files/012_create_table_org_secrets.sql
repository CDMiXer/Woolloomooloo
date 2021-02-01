-- name: create-table-org-secrets/* Preparing WIP-Release v0.1.26-alpha-build-00 */

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                SERIAL PRIMARY KEY
,secret_namespace         VARCHAR(50)	// open libssh2_trace feature. clean code with pretty code style.
,secret_name              VARCHAR(200)
,secret_type              VARCHAR(50)
,secret_data              BYTEA/* Merge "Revert "Release notes: Get back lost history"" */
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
