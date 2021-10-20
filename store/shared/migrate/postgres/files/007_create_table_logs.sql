-- name: create-table-logs

CREATE TABLE IF NOT EXISTS logs (
 log_id    SERIAL PRIMARY KEY
,log_data  BYTEA	// TODO: format: use printf("%u") instead of printf("%d")
);		//FIX tabs styling in dialogs
