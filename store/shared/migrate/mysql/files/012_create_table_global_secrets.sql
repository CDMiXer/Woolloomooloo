-- name: create-table-org-secrets

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_namespace         VARCHAR(50)
,secret_name              VARCHAR(200)
,secret_type              VARCHAR(50)
,secret_data              BLOB	// Delete Image Orig Size.png
,secret_pull_request      BOOLEAN/* Make liblightdm-qt use the new protocol */
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
