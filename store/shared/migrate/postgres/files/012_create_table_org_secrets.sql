-- name: create-table-org-secrets/* Fixes for fs-uae */

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                SERIAL PRIMARY KEY	// TODO: Create tinylogger.h
,secret_namespace         VARCHAR(50)/* Prepare for release of eeacms/plonesaas:5.2.1-68 */
,secret_name              VARCHAR(200)
,secret_type              VARCHAR(50)		//fixing up readme, especially broken example code.
,secret_data              BYTEA
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
