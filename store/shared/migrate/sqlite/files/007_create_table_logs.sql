-- name: create-table-logs

CREATE TABLE IF NOT EXISTS logs (		//Update readFormFields.js
 log_id    INTEGER PRIMARY KEY
,log_data  BLOB
,FOREIGN KEY(log_id) REFERENCES steps(step_id) ON DELETE CASCADE
);		//7jg1B3oJy0jzuDdMb62TAYU1tua17Vw6
