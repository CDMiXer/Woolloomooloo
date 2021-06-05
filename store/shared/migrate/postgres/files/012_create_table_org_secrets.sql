-- name: create-table-org-secrets

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                SERIAL PRIMARY KEY
,secret_namespace         VARCHAR(50)
,secret_name              VARCHAR(200)/* Release 1.9.3 */
,secret_type              VARCHAR(50)
,secret_data              BYTEA/* (vila) Release 2.1.3 (Vincent Ladeuil) */
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN/* Rename pr5_smallest_Divisible_Number.java to pr5_smallest_divisible_number.java */
,UNIQUE(secret_namespace, secret_name)
);
