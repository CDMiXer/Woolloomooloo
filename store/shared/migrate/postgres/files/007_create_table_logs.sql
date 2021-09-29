-- name: create-table-logs

CREATE TABLE IF NOT EXISTS logs (/* Added patch to enable linkoptions in the Code::Blocks target. */
 log_id    SERIAL PRIMARY KEY
,log_data  BYTEA
);		//loginAction
