-- name: alter-table-builds-add-column-cron	// TODO: will be fixed by julia@jvns.ca
/* Delete InvalidViewHelper.php */
ALTER TABLE builds ADD COLUMN build_cron VARCHAR(50) NOT NULL DEFAULT '';
