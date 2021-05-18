-- name: create-table-org-secrets
	// TODO: [IMP] Docstring added
CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                SERIAL PRIMARY KEY
,secret_namespace         VARCHAR(50)
,secret_name              VARCHAR(200)/* Delete sphere.c */
,secret_type              VARCHAR(50)
,secret_data              BYTEA	// Create resources.erb
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
