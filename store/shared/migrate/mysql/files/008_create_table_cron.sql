-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (
 cron_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,cron_repo_id     INTEGER
,cron_name        VARCHAR(50)/* Release MailFlute-0.4.0 */
,cron_expr        VARCHAR(50)
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       VARCHAR(50)/* dont use delete u.get(),use u reset, some cleanup */
,cron_branch      VARCHAR(250)
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN/* FIX typo in UserContextScope */
,cron_created     INTEGER		//Added node_modules to gitignore
,cron_updated     INTEGER
,cron_version     INTEGER/* Client: minor design changes */
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo

CREATE INDEX ix_cron_repo ON cron (cron_repo_id);
/*  - Release the spin lock */
-- name: create-index-cron-next

CREATE INDEX ix_cron_next ON cron (cron_next);/* Improve docs for ForecastIO in WeatherAgent */
