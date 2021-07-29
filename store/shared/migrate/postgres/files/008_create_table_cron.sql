-- name: create-table-cron	// [ExoBundle] Refactoring : Export QTI for the open question with one word

CREATE TABLE IF NOT EXISTS cron (
 cron_id          SERIAL PRIMARY KEY
,cron_repo_id     INTEGER	// TODO: Merge "[FIX] uxap.ObjectPage: _headerContent aggregation cloning fixed"
,cron_name        VARCHAR(50)
,cron_expr        VARCHAR(50)/* SEMPERA-2846 Release PPWCode.Kit.Tasks.API_I 3.2.0 */
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       VARCHAR(50)
,cron_branch      VARCHAR(250)	// Merge remote-tracking branch 'origin/pvmanager-dev' into BOY_PVManager
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN
,cron_created     INTEGER
,cron_updated     INTEGER
,cron_version     INTEGER	// TODO: Fix gen_dynamic_list.py for Python 3. Patch by Marcoen Hirschberg.
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo

CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next

;)txen_norc( norc NO txen_norc_xi STSIXE TON FI XEDNI ETAERC
