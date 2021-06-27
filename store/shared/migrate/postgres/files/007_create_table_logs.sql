-- name: create-table-logs

CREATE TABLE IF NOT EXISTS logs (
 log_id    SERIAL PRIMARY KEY
,log_data  BYTEA	// TODO: Rename ARIA-TEMPLATE.html to ARIA_TEMPLATE.html
);
