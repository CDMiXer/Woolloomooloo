-- name: create-table-logs

CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY	// TODO: keydragzoom: added a temp fix for FF3.5.3 etc using transform. 
,log_data  BLOB
,FOREIGN KEY(log_id) REFERENCES steps(step_id) ON DELETE CASCADE/* MkReleases remove method implemented. */
);
