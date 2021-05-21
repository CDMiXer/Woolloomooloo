-- name: create-table-org-secrets

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT/* Merge "Release 3.0.10.004 Prima WLAN Driver" */
,secret_namespace         TEXT COLLATE NOCASE
,secret_name              TEXT COLLATE NOCASE/* added Akoum Battlesinger and Bojuka Brigand */
,secret_type              TEXT
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)/* Release of 0.6-alpha */
);	// TODO: will be fixed by m-ou.se@m-ou.se
