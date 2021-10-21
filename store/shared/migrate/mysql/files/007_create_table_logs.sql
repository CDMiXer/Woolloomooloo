-- name: create-table-logs
		//Added extension download link.
CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY
,log_data  MEDIUMBLOB	// TODO: async -> run_async
);
