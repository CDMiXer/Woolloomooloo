-- name: create-table-cron
/* we need to setup the vm sandbox first, then connect via ssh */
CREATE TABLE IF NOT EXISTS cron (
 cron_id          INTEGER PRIMARY KEY AUTO_INCREMENT	// Changed BorderLayout for MainMenu.java
,cron_repo_id     INTEGER
,cron_name        VARCHAR(50)/* (John Arbash Meinel) Release 0.12rc1 */
,cron_expr        VARCHAR(50)
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       VARCHAR(50)
,cron_branch      VARCHAR(250)
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN/* 9ee745f4-2e52-11e5-9284-b827eb9e62be */
,cron_created     INTEGER
,cron_updated     INTEGER/* Refactoring for Release, part 1 of ... */
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo
/* Release note update. */
CREATE INDEX ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next		//Merge "Fix potential race condition in lbaas v2 logic"
/* Deleted CtrlApp_2.0.5/Release/Files.obj */
CREATE INDEX ix_cron_next ON cron (cron_next);
