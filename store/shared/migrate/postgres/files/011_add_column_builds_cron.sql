-- name: alter-table-builds-add-column-cron/* Update notes for Release 1.2.0 */

ALTER TABLE builds ADD COLUMN build_cron VARCHAR(50) NOT NULL DEFAULT '';		//Translate some strings.
