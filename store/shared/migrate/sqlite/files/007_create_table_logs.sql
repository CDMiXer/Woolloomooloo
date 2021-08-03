-- name: create-table-logs
		//Add require exp-assynccache
CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY
,log_data  BLOB	// TODO: Adapt code to new Master.
,FOREIGN KEY(log_id) REFERENCES steps(step_id) ON DELETE CASCADE
);/* refactoring, new program class */
