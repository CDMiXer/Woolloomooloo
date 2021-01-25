-- name: create-table-logs/* Release 0.32.1 */

CREATE TABLE IF NOT EXISTS logs (	// TODO: will be fixed by martin2cai@hotmail.com
 log_id    SERIAL PRIMARY KEY
,log_data  BYTEA
);
