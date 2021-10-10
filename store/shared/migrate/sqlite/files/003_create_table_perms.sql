-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (		//Added process.stop function which halts the QMiner instance.
 perm_user_id  INTEGER
,perm_repo_uid TEXT
,perm_read     BOOLEAN
,perm_write    BOOLEAN	// TODO: Style improved.
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);
/* bugfix to sass format. */
-- name: create-index-perms-user	// 555eca8c-5216-11e5-99c0-6c40088e03e4

CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);/* IHTSDO Release 4.5.51 */
		//Added a patch method.
-- name: create-index-perms-repo	// TODO: will be fixed by aeongrp@outlook.com

CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);
