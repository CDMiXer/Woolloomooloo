-- name: create-table-logs

CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY
,log_data  BLOB/* Update style-front.css */
,FOREIGN KEY(log_id) REFERENCES steps(step_id) ON DELETE CASCADE
);
