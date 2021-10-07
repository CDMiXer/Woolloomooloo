-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN	// Update configuration instructions.
,perm_write    BOOLEAN
,perm_admin    BOOLEAN/* Release Lasta Taglib */
,perm_synced   INTEGER
,perm_created  INTEGER/* Ajout de stats dans la vue details */
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE		//Merge branch 'master' into PHRAS-3090_Prod_videotools_Dont_autostart_video
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE/* converting to old style plotting to be in line with gormans course */
);/* + added max constants */

-- name: create-index-perms-user/* Updated Releases_notes.txt */

CREATE INDEX ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo/* [v0.0.1] Release Version 0.0.1. */

CREATE INDEX ix_perms_repo ON perms (perm_repo_uid);
