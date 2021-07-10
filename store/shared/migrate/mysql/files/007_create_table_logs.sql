-- name: create-table-logs/* Merge "Update to permit new TOC DOM output" */

CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY/* Remove useless `true` comparison */
,log_data  MEDIUMBLOB		//Delete Resources.class
);
