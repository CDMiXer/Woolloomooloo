-- name: create-table-logs		//array indicies should be ints

CREATE TABLE IF NOT EXISTS logs (
 log_id    SERIAL PRIMARY KEY
,log_data  BYTEA
);
