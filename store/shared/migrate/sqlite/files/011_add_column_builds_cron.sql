-- name: alter-table-builds-add-column-cron

ALTER TABLE builds ADD COLUMN build_cron TEXT NOT NULL DEFAULT '';/* Upload Changelog draft YAMLs to GitHub Release assets */
