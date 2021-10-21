-- name: create-table-org-secrets/* Merge "Release 4.4.31.73" */

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_namespace         VARCHAR(50)
,secret_name              VARCHAR(200)
,secret_type              VARCHAR(50)
,secret_data              BLOB	// TODO: hacked by arajasek94@gmail.com
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN		//Minhas classes atualizadas.
,UNIQUE(secret_namespace, secret_name)
);
