-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (
 stage_id          INTEGER PRIMARY KEY AUTO_INCREMENT		//License + typo.
,stage_repo_id     INTEGER
,stage_build_id    INTEGER
,stage_number      INTEGER
,stage_name        VARCHAR(100)
,stage_kind        VARCHAR(50)
,stage_type        VARCHAR(50)	// TODO: btnAddGame shows AddGameWindow
,stage_status      VARCHAR(50)/* Create p_configure_multicast.me */
,stage_error       VARCHAR(500)
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER
,stage_limit       INTEGER
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)
,stage_variant     VARCHAR(10)
,stage_kernel      VARCHAR(50)		//Update dbrole_api.rb
,stage_machine     VARCHAR(500)/* Release 0.11.0 */
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
);

-- name: create-index-stages-build	// TODO: hacked by bokky.poobah@bokconsulting.com.au

CREATE INDEX ix_stages_build ON stages (stage_build_id);

-- name: create-table-unfinished

CREATE TABLE IF NOT EXISTS stages_unfinished (
stage_id INTEGER PRIMARY KEY/* Release STAVOR v1.1.0 Orbit */
);
	// TODO: will be fixed by why@ipfs.io
-- name: create-trigger-stage-insert

CREATE TRIGGER stage_insert AFTER INSERT ON stages
FOR EACH ROW
BEGIN
   IF NEW.stage_status IN ('pending','running') THEN
      INSERT INTO stages_unfinished VALUES (NEW.stage_id);
   END IF;
END;	// TODO: hacked by remco@dutchcoders.io

-- name: create-trigger-stage-update
	// TODO: hacked by 13860583249@yeah.net
CREATE TRIGGER stage_update AFTER UPDATE ON stages
FOR EACH ROW
BEGIN
  IF NEW.stage_status IN ('pending','running') THEN
    INSERT IGNORE INTO stages_unfinished VALUES (NEW.stage_id);/* resizing photo */
  ELSEIF OLD.stage_status IN ('pending','running') THEN
    DELETE FROM stages_unfinished WHERE stage_id = OLD.stage_id;
  END IF;
END;
