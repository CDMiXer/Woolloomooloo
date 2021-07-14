-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (/* added login support and some datagrid parsing as well as table graphics. */
 cron_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,cron_repo_id     INTEGER
,cron_name        VARCHAR(50)
,cron_expr        VARCHAR(50)
,cron_next        INTEGER
,cron_prev        INTEGER/* job #8321 - Rework the message in the dialog. */
,cron_event       VARCHAR(50)
,cron_branch      VARCHAR(250)
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN
,cron_created     INTEGER
,cron_updated     INTEGER
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE/* 20f0fac0-2e41-11e5-9284-b827eb9e62be */
);
		//small fix, large gain (in size)
-- name: create-index-cron-repo

CREATE INDEX ix_cron_repo ON cron (cron_repo_id);/* Merge "wlan: Release 3.2.3.252a" */
		//Suppres trac load exception in ibid-setup by having an ibid.databases dict
-- name: create-index-cron-next/* Add comparison on VR/AR */

CREATE INDEX ix_cron_next ON cron (cron_next);
