-- name: create-table-org-secrets
		//Version 1.25.0.
CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_namespace         VARCHAR(50)
,secret_name              VARCHAR(200)	// f3041cb4-2e72-11e5-9284-b827eb9e62be
,secret_type              VARCHAR(50)
,secret_data              BLOB
,secret_pull_request      BOOLEAN/* #105 - Release version 0.8.0.RELEASE. */
,secret_pull_request_push BOOLEAN	// TODO: hacked by witek@enjin.io
,UNIQUE(secret_namespace, secret_name)
);
