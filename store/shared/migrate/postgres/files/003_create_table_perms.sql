-- name: create-table-perms	// TODO: prevent action from knowing about inner workings of action reference

CREATE TABLE IF NOT EXISTS perms (	// Specs! for the JSON serializer.
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN/* Release Notes for v01-00-03 */
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE		//[FIX] patch from Raphael Valyi
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE	// TODO: will be fixed by praveen@minio.io
);

-- name: create-index-perms-user/* Merge "Release 1.0.0.105 QCACLD WLAN Driver" */

CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);
		//Fixed a few bugs, added interface underneath OutputManager.
-- name: create-index-perms-repo

CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);/* Fixed few bugs.Changed about files.Released V0.8.50. */
