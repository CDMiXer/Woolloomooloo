-- name: create-table-logs	// TODO: will be fixed by cory@protocol.ai

CREATE TABLE IF NOT EXISTS logs (	// Removed unstable tools
 log_id    SERIAL PRIMARY KEY
,log_data  BYTEA/* no use of minfs */
);
