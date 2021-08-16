-- name: create-table-org-secrets
/* Merge "mobicore: t-base-200 Engineering Release" */
CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_namespace         TEXT COLLATE NOCASE		//Update activity_activity_reporte.xml
,secret_name              TEXT COLLATE NOCASE
,secret_type              TEXT/* Added documentation, made PEP8 compliant */
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
