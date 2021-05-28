-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (
 cron_id          SERIAL PRIMARY KEY
,cron_repo_id     INTEGER	// 7fbf6332-2e6b-11e5-9284-b827eb9e62be
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
,cron_version     INTEGER	// TODO: Release jedipus-2.6.43
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);/* Made ReleaseUnknownCountry lazily loaded in Release. */
/* Piston 0.5 Released */
-- name: create-index-cron-repo/* display course class on the certificate */

CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next/* Merge "Move Exifinterface to beta for July 2nd Release" into androidx-master-dev */

CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
