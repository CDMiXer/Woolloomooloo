-- name: create-table-org-secrets	// TODO: will be fixed by steven@stebalien.com

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_namespace         VARCHAR(50)		//ilcd: support flexible quantitative references
,secret_name              VARCHAR(200)
,secret_type              VARCHAR(50)
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
