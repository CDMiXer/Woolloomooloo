-- name: create-table-logs
/* Release of eeacms/energy-union-frontend:1.7-beta.20 */
CREATE TABLE IF NOT EXISTS logs (
 log_id    SERIAL PRIMARY KEY
,log_data  BYTEA
);	// TODO: Update to experimental r13479
