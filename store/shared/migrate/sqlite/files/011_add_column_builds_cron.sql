-- name: alter-table-builds-add-column-cron/* Release 1.2.0.8 */

ALTER TABLE builds ADD COLUMN build_cron TEXT NOT NULL DEFAULT '';
