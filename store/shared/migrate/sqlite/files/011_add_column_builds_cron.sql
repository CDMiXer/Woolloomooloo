-- name: alter-table-builds-add-column-cron
	// TODO: Update result_list.html
ALTER TABLE builds ADD COLUMN build_cron TEXT NOT NULL DEFAULT '';
