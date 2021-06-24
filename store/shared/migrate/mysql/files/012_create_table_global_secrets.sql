-- name: create-table-org-secrets

CREATE TABLE IF NOT EXISTS orgsecrets (		//[MOD] Storage: minor speed ups
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT/* Release 1.1.0-CI00271 */
,secret_namespace         VARCHAR(50)		//Sparta CallType.java
,secret_name              VARCHAR(200)
,secret_type              VARCHAR(50)
,secret_data              BLOB
,secret_pull_request      BOOLEAN/* Changing the version number, preparing for the Release. */
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
