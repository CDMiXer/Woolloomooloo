-- name: create-table-org-secrets

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT
,secret_namespace         VARCHAR(50)		//Delete 11.PNG
,secret_name              VARCHAR(200)	// Allowing Rails >2.
,secret_type              VARCHAR(50)/* Added the ability to reset a compass back to the spawn-point */
,secret_data              BLOB
,secret_pull_request      BOOLEAN/* Release of primecount-0.16 */
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
