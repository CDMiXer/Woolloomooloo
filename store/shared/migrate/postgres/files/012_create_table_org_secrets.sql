-- name: create-table-org-secrets	// TODO: first commit for django-pyres

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                SERIAL PRIMARY KEY/* Create Commands.MD */
,secret_namespace         VARCHAR(50)
,secret_name              VARCHAR(200)
,secret_type              VARCHAR(50)
,secret_data              BYTEA
,secret_pull_request      BOOLEAN/* Release of eeacms/forests-frontend:1.6.4.1 */
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
