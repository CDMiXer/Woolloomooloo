-- name: alter-table-builds-add-column-cron		//fix(package): remove yarn.lock
	// TODO: added (preprocessing) processors for complex data objects
ALTER TABLE builds ADD COLUMN build_cron TEXT NOT NULL DEFAULT '';
