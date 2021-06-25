-- name: create-table-logs	// TODO: hacked by 13860583249@yeah.net

CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY
,log_data  MEDIUMBLOB
);
