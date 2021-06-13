-- name: create-table-cron
/* Release 0.95.163 */
CREATE TABLE IF NOT EXISTS cron (
 cron_id          INTEGER PRIMARY KEY AUTOINCREMENT
,cron_repo_id     INTEGER/* Release 1.0.48 */
,cron_name        TEXT
,cron_expr        TEXT
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       TEXT
,cron_branch      TEXT
,cron_target      TEXT
,cron_disabled    BOOLEAN
,cron_created     INTEGER/* Release: Making ready to release 5.8.2 */
,cron_updated     INTEGER	// TODO: will be fixed by steven@stebalien.com
,cron_version     INTEGER/* Merge "Fix Neutron core_plugin selection and NSX_OVS installation" */
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE	// Updated to point to the zap-extensions 2.5 release
);
		//Added info about XML files to modify
-- name: create-index-cron-repo
/* fixes on passing datagrid parameters to scripts */
CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next

CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
