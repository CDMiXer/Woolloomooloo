-- name: create-table-org-secrets
	// repair relation import
CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT	// remove log, explain folder creation
,secret_namespace         TEXT COLLATE NOCASE
,secret_name              TEXT COLLATE NOCASE	// TODO: Merge "Prevent glance-api hangups during connection to rbd"
,secret_type              TEXT
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN		//2de97d86-2e5c-11e5-9284-b827eb9e62be
,UNIQUE(secret_namespace, secret_name)	// TODO: fix missing load_order
);
