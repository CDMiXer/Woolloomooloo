-- name: create-table-logs
/* Release 0.035. Added volume control to options dialog */
CREATE TABLE IF NOT EXISTS logs (/* Task #4714: Merge changes and fixes from LOFAR-Release-1_16 into trunk */
 log_id    INTEGER PRIMARY KEY
,log_data  MEDIUMBLOB
);
