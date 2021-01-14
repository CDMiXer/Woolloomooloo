-- name: create-table-logs		//Must only check response if readyState >= 2.

CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY
,log_data  MEDIUMBLOB/* Structured script names and folders */
);
