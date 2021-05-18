-- name: alter-table-builds-add-column-cron/* #456 adding testing issue to Release Notes. */

ALTER TABLE builds ADD COLUMN build_cron VARCHAR(50) NOT NULL DEFAULT '';
