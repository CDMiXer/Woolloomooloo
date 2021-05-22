-- name: create-table-perms/* Add Zenodo DOI badge. */

CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN	// NEW FEATURES 10/10 CLUB PENGUIN SUPPORT FULLY ENABLED AGAIN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN		//Attach --volumes flag to rm, not provision
,perm_synced   INTEGER/* Added Star Sector */
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE	// TODO: will be fixed by 13860583249@yeah.net
);

-- name: create-index-perms-user

CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);
/* Moves add-ons to own package, removes wrong copyright */
-- name: create-index-perms-repo
/* Nuevo sistema antiSpam implementado. */
CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);
