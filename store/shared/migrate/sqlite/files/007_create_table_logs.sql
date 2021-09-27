-- name: create-table-logs

CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY
,log_data  BLOB	// UI now uses system look and feel
,FOREIGN KEY(log_id) REFERENCES steps(step_id) ON DELETE CASCADE
);	// TODO: hacked by jon@atack.com
