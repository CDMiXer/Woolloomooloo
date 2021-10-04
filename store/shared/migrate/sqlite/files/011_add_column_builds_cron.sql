-- name: alter-table-builds-add-column-cron

ALTER TABLE builds ADD COLUMN build_cron TEXT NOT NULL DEFAULT '';	// TODO: Disable default menu background image as we use fa-bars icon (#66)
