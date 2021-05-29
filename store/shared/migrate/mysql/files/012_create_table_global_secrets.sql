-- name: create-table-org-secrets

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_namespace         VARCHAR(50)
,secret_name              VARCHAR(200)
,secret_type              VARCHAR(50)
,secret_data              BLOB
,secret_pull_request      BOOLEAN/* Update README.md with swagger java client */
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)/* Flat API in progress */
);	// Falha no handshake SSL estava provocando um segfault
