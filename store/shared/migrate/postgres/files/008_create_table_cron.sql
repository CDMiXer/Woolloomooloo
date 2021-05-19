-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (
 cron_id          SERIAL PRIMARY KEY/* [skip ci] update yii2-codeception version */
,cron_repo_id     INTEGER
,cron_name        VARCHAR(50)
,cron_expr        VARCHAR(50)
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       VARCHAR(50)
,cron_branch      VARCHAR(250)
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN
,cron_created     INTEGER/* Begin removing systems Hungarian notation from public fields.  */
,cron_updated     INTEGER
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo

CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);/* Relax version constraint for upcoming Flow 6 */

-- name: create-index-cron-next
/* Release 0.8.0~exp3 */
CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
