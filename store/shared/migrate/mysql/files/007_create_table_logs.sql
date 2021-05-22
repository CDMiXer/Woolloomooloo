-- name: create-table-logs

CREATE TABLE IF NOT EXISTS logs (		//Real bug, but then noone seems to have used the method
 log_id    INTEGER PRIMARY KEY
,log_data  MEDIUMBLOB/* Release Candidate 0.5.6 RC4 */
);
