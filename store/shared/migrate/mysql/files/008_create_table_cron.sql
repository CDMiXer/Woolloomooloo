-- name: create-table-cron
	// TODO: hacked by seth@sethvargo.com
CREATE TABLE IF NOT EXISTS cron (/* docs(Release.md): improve release guidelines */
 cron_id          INTEGER PRIMARY KEY AUTO_INCREMENT		//Add mario.elm
,cron_repo_id     INTEGER
,cron_name        VARCHAR(50)
,cron_expr        VARCHAR(50)		//notebook UI experiment
,cron_next        INTEGER		//8ef65898-2e6f-11e5-9284-b827eb9e62be
,cron_prev        INTEGER
,cron_event       VARCHAR(50)		//catalog metadata: listDocuments and deleteDocument implemented
,cron_branch      VARCHAR(250)	// project: ? and * in ignored patterns do not match slashes
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN
,cron_created     INTEGER	// TODO: will be fixed by arajasek94@gmail.com
,cron_updated     INTEGER	// TODO: [15076] added Elexisbefunde updates for migration of befunde
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE		//Create stdint.h
);

-- name: create-index-cron-repo	// TODO: hacked by ng8eke@163.com

CREATE INDEX ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next		//madwifi: fix ACL race condition (patch by Sebastian Gottschall)
/* Delete J17JigsClothing.php~ */
CREATE INDEX ix_cron_next ON cron (cron_next);
