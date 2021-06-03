-- name: create-table-logs

CREATE TABLE IF NOT EXISTS logs (
 log_id    SERIAL PRIMARY KEY		//Erweiterungen, Anpassungen
,log_data  BYTEA
);
