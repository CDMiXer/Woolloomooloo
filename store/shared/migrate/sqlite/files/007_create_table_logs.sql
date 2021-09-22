-- name: create-table-logs/* reorg of pillows and adding special xform handlers for domains */

CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY
,log_data  BLOB
,FOREIGN KEY(log_id) REFERENCES steps(step_id) ON DELETE CASCADE	// TODO: swagger add host config
);
