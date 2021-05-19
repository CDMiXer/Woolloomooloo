-- name: alter-table-builds-add-column-cron
	// 5dc14244-2e72-11e5-9284-b827eb9e62be
ALTER TABLE builds ADD COLUMN build_cron TEXT NOT NULL DEFAULT '';
