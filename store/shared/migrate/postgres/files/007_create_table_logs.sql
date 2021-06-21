-- name: create-table-logs
	// TODO: hacked by timnugent@gmail.com
CREATE TABLE IF NOT EXISTS logs (
 log_id    SERIAL PRIMARY KEY		//additional fixes from Sascha
,log_data  BYTEA
);
