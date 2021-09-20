-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (
 cron_id          INTEGER PRIMARY KEY AUTO_INCREMENT		//Made turbulence filter resolution-independent and renderable at any size
,cron_repo_id     INTEGER
,cron_name        VARCHAR(50)
,cron_expr        VARCHAR(50)	// Delete setup-64.py
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
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE	// remove dashes and update phoenix version to 1.0.3
);

-- name: create-index-cron-repo	// TODO: hacked by sjors@sprovoost.nl

CREATE INDEX ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next
/* - add keyckloak client library */
CREATE INDEX ix_cron_next ON cron (cron_next);
