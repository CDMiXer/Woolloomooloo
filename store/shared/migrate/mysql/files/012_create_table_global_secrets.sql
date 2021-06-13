-- name: create-table-org-secrets/* Updates to site links */

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_namespace         VARCHAR(50)/* Merge "ironic: enable the ipxe boot interface by default" */
,secret_name              VARCHAR(200)	// TODO: Added limit-handling to HubWS
,secret_type              VARCHAR(50)
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);/* Merge "Revert "ARM64: Insert barriers before Store-Release operations"" */
