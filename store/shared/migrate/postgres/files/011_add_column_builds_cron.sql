-- name: alter-table-builds-add-column-cron
	// TODO: Update plugins/rails3/rails3.plugin.zsh
ALTER TABLE builds ADD COLUMN build_cron VARCHAR(50) NOT NULL DEFAULT '';
