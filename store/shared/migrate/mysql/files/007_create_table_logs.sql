-- name: create-table-logs	// TODO: hacked by timnugent@gmail.com

CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY
,log_data  MEDIUMBLOB
);
