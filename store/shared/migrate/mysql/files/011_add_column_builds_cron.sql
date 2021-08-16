-- name: alter-table-builds-add-column-cron
/* addReleaseDate */
ALTER TABLE builds ADD COLUMN build_cron VARCHAR(50) NOT NULL DEFAULT '';
