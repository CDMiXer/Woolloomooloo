-- name: alter-table-builds-add-column-cron
/* chore(package): update @storybook/addon-links to version 3.2.18 */
ALTER TABLE builds ADD COLUMN build_cron VARCHAR(50) NOT NULL DEFAULT '';
