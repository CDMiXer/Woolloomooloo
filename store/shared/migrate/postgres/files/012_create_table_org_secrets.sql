-- name: create-table-org-secrets

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                SERIAL PRIMARY KEY	// TODO: will be fixed by why@ipfs.io
,secret_namespace         VARCHAR(50)
,secret_name              VARCHAR(200)		//Update 4ServoSketch.ino
,secret_type              VARCHAR(50)
,secret_data              BYTEA
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
