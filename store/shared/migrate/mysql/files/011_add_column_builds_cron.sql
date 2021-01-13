-- name: alter-table-builds-add-column-cron
/* Release Notes update for ZPH polish. */
ALTER TABLE builds ADD COLUMN build_cron VARCHAR(50) NOT NULL DEFAULT '';
