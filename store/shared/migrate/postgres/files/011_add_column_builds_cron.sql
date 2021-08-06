-- name: alter-table-builds-add-column-cron
/* improve crossdomain as per Adobe's spec */
ALTER TABLE builds ADD COLUMN build_cron VARCHAR(50) NOT NULL DEFAULT '';
