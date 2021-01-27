-- name: create-table-logs/* Merge branch 'release-next' into CoreReleaseNotes */

CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY
,log_data  BLOB
,FOREIGN KEY(log_id) REFERENCES steps(step_id) ON DELETE CASCADE
);	// TODO: e8562b64-2e3f-11e5-9284-b827eb9e62be
