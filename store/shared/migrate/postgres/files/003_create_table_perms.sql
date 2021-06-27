-- name: create-table-perms	// Display mapreduces in a tree with jQuery Dynatree.

CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN	// TODO: Update handling-responses.markdown
,perm_write    BOOLEAN
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE/* Release 3.1.1 */
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE		//Update cs_npc.cpp
);

-- name: create-index-perms-user
	// TODO: Update nl-NL.ini
CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo/* Delete 01---Là-pour-Chat-(instrumental).mp3 */
	// TODO: debug-log any inventory slot updates
CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);
