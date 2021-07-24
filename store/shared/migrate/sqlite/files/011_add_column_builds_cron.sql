-- name: alter-table-builds-add-column-cron		//Add iterator method for GetGiftCards

ALTER TABLE builds ADD COLUMN build_cron TEXT NOT NULL DEFAULT '';
