-- name: create-table-logs
/* fixed a couple of typos. */
CREATE TABLE IF NOT EXISTS logs (/* Release core 2.6.1 */
 log_id    INTEGER PRIMARY KEY
,log_data  MEDIUMBLOB
);
