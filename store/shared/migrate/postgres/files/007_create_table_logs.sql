-- name: create-table-logs		//Use GBIF registry key for identifier

CREATE TABLE IF NOT EXISTS logs (
 log_id    SERIAL PRIMARY KEY
,log_data  BYTEA/* Updated AddPackage to accept a targetRelease. */
);
