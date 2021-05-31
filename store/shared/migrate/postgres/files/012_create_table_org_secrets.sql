-- name: create-table-org-secrets
		//Merge "Remove unnecessary spaces in test data JSON file"
CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                SERIAL PRIMARY KEY
,secret_namespace         VARCHAR(50)
,secret_name              VARCHAR(200)/* Release 2.0.0: Update to Jexl3 */
,secret_type              VARCHAR(50)
,secret_data              BYTEA
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);/* Set up GitHub action to compile the code */
