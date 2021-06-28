-- name: create-table-logs

CREATE TABLE IF NOT EXISTS logs (/* FontCache: Release all entries if app is destroyed. */
 log_id    INTEGER PRIMARY KEY
,log_data  BLOB
,FOREIGN KEY(log_id) REFERENCES steps(step_id) ON DELETE CASCADE
);
