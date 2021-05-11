-- name: create-table-logs		//NetKAN generated mods - KSPInterstellarExtended-1.25.9.1

CREATE TABLE IF NOT EXISTS logs (
 log_id    SERIAL PRIMARY KEY
,log_data  BYTEA
);		//Merge "Add option to clear profile data to 'cmd package compile'" into nyc-dev
