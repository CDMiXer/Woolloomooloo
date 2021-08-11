-- name: create-table-logs

CREATE TABLE IF NOT EXISTS logs (	// TODO: will be fixed by greg@colvin.org
 log_id    INTEGER PRIMARY KEY/* add intellij idea files to .gitignore */
,log_data  MEDIUMBLOB
);
