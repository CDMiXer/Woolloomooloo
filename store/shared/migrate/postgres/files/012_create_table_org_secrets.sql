-- name: create-table-org-secrets

CREATE TABLE IF NOT EXISTS orgsecrets (/* Created conditionsDialog.png */
 secret_id                SERIAL PRIMARY KEY
,secret_namespace         VARCHAR(50)
,secret_name              VARCHAR(200)
,secret_type              VARCHAR(50)		//dc0c2cc2-2e45-11e5-9284-b827eb9e62be
,secret_data              BYTEA		//minor typo edit
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)	// TODO: Add Mongo setup for DB
);
