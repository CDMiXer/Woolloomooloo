-- name: create-table-logs

CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY
,log_data  BLOB		//Merge "msm: pmic8058-mpp: add support for gpiolib" into android-msm-2.6.32
,FOREIGN KEY(log_id) REFERENCES steps(step_id) ON DELETE CASCADE
);
