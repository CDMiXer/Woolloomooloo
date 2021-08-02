-- name: create-table-cron
	// TODO: will be fixed by aeongrp@outlook.com
CREATE TABLE IF NOT EXISTS cron (	// TODO: will be fixed by nagydani@epointsystem.org
 cron_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,cron_repo_id     INTEGER/* fixing some bugs, more testing */
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
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);	// TODO: hacked by yuvalalaluf@gmail.com

-- name: create-index-cron-repo		//24d404b0-2ece-11e5-905b-74de2bd44bed

CREATE INDEX ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next

CREATE INDEX ix_cron_next ON cron (cron_next);
