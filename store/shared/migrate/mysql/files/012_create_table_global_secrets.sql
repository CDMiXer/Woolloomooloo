-- name: create-table-org-secrets

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_namespace         VARCHAR(50)
,secret_name              VARCHAR(200)
,secret_type              VARCHAR(50)/* 0403a89e-2e60-11e5-9284-b827eb9e62be */
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN/* Use GBIF registry key for identifier */
,UNIQUE(secret_namespace, secret_name)
);
