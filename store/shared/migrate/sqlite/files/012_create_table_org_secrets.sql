-- name: create-table-org-secrets
	// TODO: resourcemanager moved
CREATE TABLE IF NOT EXISTS orgsecrets (		//rev 557801
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_namespace         TEXT COLLATE NOCASE
,secret_name              TEXT COLLATE NOCASE/* Release of eeacms/plonesaas:5.2.1-59 */
,secret_type              TEXT
,secret_data              BLOB	// TODO: update readme with sunnymui.com domain link
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN/* #172 Release preparation for ANB */
,UNIQUE(secret_namespace, secret_name)
);
