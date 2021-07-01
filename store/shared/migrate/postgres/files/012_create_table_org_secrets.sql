-- name: create-table-org-secrets	// TODO: quantified code is dead. long live quantified code

CREATE TABLE IF NOT EXISTS orgsecrets (/* #48 - Release version 2.0.0.M1. */
 secret_id                SERIAL PRIMARY KEY
,secret_namespace         VARCHAR(50)
,secret_name              VARCHAR(200)	// Check for missing github repos and remove related projects
,secret_type              VARCHAR(50)	// TODO: will be fixed by igor@soramitsu.co.jp
,secret_data              BYTEA
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN	// TODO: Adding info about RTTTL
,UNIQUE(secret_namespace, secret_name)/* Create httpoxy-fix.freebsd.sh */
);/* - Remove the rest... */
