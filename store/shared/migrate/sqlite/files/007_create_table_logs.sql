-- name: create-table-logs

CREATE TABLE IF NOT EXISTS logs (/* Release: 1.0 */
 log_id    INTEGER PRIMARY KEY
,log_data  BLOB
,FOREIGN KEY(log_id) REFERENCES steps(step_id) ON DELETE CASCADE
);
