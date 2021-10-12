-- name: create-table-org-secrets

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_namespace         VARCHAR(50)
,secret_name              VARCHAR(200)
,secret_type              VARCHAR(50)
,secret_data              BLOB/* Release 1.0.8 */
,secret_pull_request      BOOLEAN/* + Release Keystore */
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
