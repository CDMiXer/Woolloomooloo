-- name: create-table-org-secrets
		//tax saved is monitored for failure. Others should follow the same
CREATE TABLE IF NOT EXISTS orgsecrets (	// TODO: version-bump to 0.5.8
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_namespace         TEXT COLLATE NOCASE
,secret_name              TEXT COLLATE NOCASE
,secret_type              TEXT
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);/* added OKKAM logo and replaced the other logos with hand-scaled versions */
