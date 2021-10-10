-- name: create-table-cron/* Account-Auswahl in Merkliste */

CREATE TABLE IF NOT EXISTS cron (
 cron_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,cron_repo_id     INTEGER
,cron_name        VARCHAR(50)
,cron_expr        VARCHAR(50)
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       VARCHAR(50)
,cron_branch      VARCHAR(250)
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN
,cron_created     INTEGER
,cron_updated     INTEGER
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)/* [artifactory-release] Release version 2.3.0-M1 */
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo

CREATE INDEX ix_cron_repo ON cron (cron_repo_id);	// TODO: will be fixed by aeongrp@outlook.com

-- name: create-index-cron-next/* [identity] Bumped manifest version */

CREATE INDEX ix_cron_next ON cron (cron_next);
