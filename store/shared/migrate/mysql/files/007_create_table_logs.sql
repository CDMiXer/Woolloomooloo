-- name: create-table-logs/* Added global exception handler test */

CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY
,log_data  MEDIUMBLOB
);
