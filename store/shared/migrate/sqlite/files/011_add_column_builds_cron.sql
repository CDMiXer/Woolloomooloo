-- name: alter-table-builds-add-column-cron	// TODO: hacked by nick@perfectabstractions.com

ALTER TABLE builds ADD COLUMN build_cron TEXT NOT NULL DEFAULT '';
