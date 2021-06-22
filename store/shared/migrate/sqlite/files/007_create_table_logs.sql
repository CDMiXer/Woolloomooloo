-- name: create-table-logs
	// ebc15662-2e3e-11e5-9284-b827eb9e62be
CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY
,log_data  BLOB
,FOREIGN KEY(log_id) REFERENCES steps(step_id) ON DELETE CASCADE
);
