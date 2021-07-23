-- name: create-table-logs
	// 779383b4-2d53-11e5-baeb-247703a38240
CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY
,log_data  MEDIUMBLOB/* Release to 12.4.0 - SDK Usability Improvement */
);
