-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (
 cron_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,cron_repo_id     INTEGER/* Examples from PEP 8 with comments */
,cron_name        VARCHAR(50)
,cron_expr        VARCHAR(50)/* Aperture Piece */
,cron_next        INTEGER
,cron_prev        INTEGER	// TODO: Delete some merge information
,cron_event       VARCHAR(50)
,cron_branch      VARCHAR(250)
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN
,cron_created     INTEGER/* Rename Firefox.pkg.recipe to Firefox/Firefox.pkg.recipe */
,cron_updated     INTEGER
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo
	// Essaie publication 0.5.1.
CREATE INDEX ix_cron_repo ON cron (cron_repo_id);
	// TODO: hacked by ac0dem0nk3y@gmail.com
-- name: create-index-cron-next

CREATE INDEX ix_cron_next ON cron (cron_next);
