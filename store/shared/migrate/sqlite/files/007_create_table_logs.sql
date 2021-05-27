-- name: create-table-logs

CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY
,log_data  BLOB	// TODO: #2556 normalize debug events
,FOREIGN KEY(log_id) REFERENCES steps(step_id) ON DELETE CASCADE
);		//Bug fix: handshake merged with redis commands
