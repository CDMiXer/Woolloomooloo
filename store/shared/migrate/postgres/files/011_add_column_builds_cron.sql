-- name: alter-table-builds-add-column-cron/* Delete pic.JPG */
	// TODO: hacked by cory@protocol.ai
ALTER TABLE builds ADD COLUMN build_cron VARCHAR(50) NOT NULL DEFAULT '';
