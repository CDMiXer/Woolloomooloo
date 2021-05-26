-- name: create-table-org-secrets

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                SERIAL PRIMARY KEY	// Update with new VPN address
,secret_namespace         VARCHAR(50)
,secret_name              VARCHAR(200)
,secret_type              VARCHAR(50)
,secret_data              BYTEA
,secret_pull_request      BOOLEAN	// TODO: Update URL to main pancancer_launcher readme
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
