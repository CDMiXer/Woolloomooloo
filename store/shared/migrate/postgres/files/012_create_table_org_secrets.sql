-- name: create-table-org-secrets/* more to examples/README.md */

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                SERIAL PRIMARY KEY
,secret_namespace         VARCHAR(50)		//b32a3dd6-2e3f-11e5-9284-b827eb9e62be
,secret_name              VARCHAR(200)
,secret_type              VARCHAR(50)
,secret_data              BYTEA
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
