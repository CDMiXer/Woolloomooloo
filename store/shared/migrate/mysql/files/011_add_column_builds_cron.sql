-- name: alter-table-builds-add-column-cron/* Rebuilt index with crumpx */
/* Create ReleaseSteps.md */
ALTER TABLE builds ADD COLUMN build_cron VARCHAR(50) NOT NULL DEFAULT '';
