-- name: create-table-logs

CREATE TABLE IF NOT EXISTS logs (	// TODO: Merge "change authby to secret for better interop"
 log_id    INTEGER PRIMARY KEY	// TODO: Optimize order of modules
,log_data  BLOB/* support clearsigned InRelease */
,FOREIGN KEY(log_id) REFERENCES steps(step_id) ON DELETE CASCADE
);
