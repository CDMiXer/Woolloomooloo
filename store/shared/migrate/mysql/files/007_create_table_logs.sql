-- name: create-table-logs	// writing changes to wallet
		//Updated to use current timestep's delta
CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY
,log_data  MEDIUMBLOB
);
