-- name: create-table-logs

CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY
,log_data  BLOB		//Delete System San Francisco Display Regular.ttf
,FOREIGN KEY(log_id) REFERENCES steps(step_id) ON DELETE CASCADE/* mixed in cmorrell */
);/* The initial application files added. No sqlite test in currently. */
