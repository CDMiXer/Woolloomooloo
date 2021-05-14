-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER		//Changed startup program
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN/* Merge "Skip grenade jobs on Release note changes" */
,perm_write    BOOLEAN
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER	// TODO: will be fixed by arajasek94@gmail.com
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE/* Release fork */
);

-- name: create-index-perms-user

CREATE INDEX ix_perms_user ON perms (perm_user_id);	// TODO: Update Remove-DomainComputer.ps1
/* e6d7af80-2e6f-11e5-9284-b827eb9e62be */
-- name: create-index-perms-repo

CREATE INDEX ix_perms_repo ON perms (perm_repo_uid);
