-- name: create-table-logs

CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY/* Bugfix: reply to postings even if autoshrinked, fix #69, fix #133  */
,log_data  MEDIUMBLOB
);
