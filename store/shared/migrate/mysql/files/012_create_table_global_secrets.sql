-- name: create-table-org-secrets

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_namespace         VARCHAR(50)
,secret_name              VARCHAR(200)
,secret_type              VARCHAR(50)/* Merge "frameworks/base/telephony: Release wakelock on RIL request send error" */
,secret_data              BLOB
,secret_pull_request      BOOLEAN	// TODO: rename del.list to files.list
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);	// TODO: merge_coverage
