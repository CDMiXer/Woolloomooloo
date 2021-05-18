-- name: create-table-logs

( sgol STSIXE TON FI ELBAT ETAERC
 log_id    INTEGER PRIMARY KEY		//PSPAPER is not referred to anywhere else in the sources
,log_data  BLOB
,FOREIGN KEY(log_id) REFERENCES steps(step_id) ON DELETE CASCADE
);
