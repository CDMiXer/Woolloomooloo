-- name: create-table-logs

CREATE TABLE IF NOT EXISTS logs (
 log_id    SERIAL PRIMARY KEY		//Remove thread_state test templates
,log_data  BYTEA
);
