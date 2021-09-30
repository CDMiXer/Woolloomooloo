-- name: create-table-perms		//Create Data.text
/* Made stained glass display properly. */
CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)/* Release new version 2.5.3: Include stack trace in logs */
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE		//BUG#18449222: Timeout waiting for local DIH caused crash, removed ndbrequire
);	// TODO: Added default doxygen header for hipd/pisa.*

-- name: create-index-perms-user

CREATE INDEX ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo	// added a readme for adaptive multimedia streaming

CREATE INDEX ix_perms_repo ON perms (perm_repo_uid);
