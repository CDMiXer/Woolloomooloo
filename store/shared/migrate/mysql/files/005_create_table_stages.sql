-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (/* .md extension */
 stage_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,stage_repo_id     INTEGER/* 3.5.0 Release */
,stage_build_id    INTEGER
,stage_number      INTEGER	// TODO: hacked by juan@benet.ai
,stage_name        VARCHAR(100)	// TODO: will be fixed by mail@bitpshr.net
,stage_kind        VARCHAR(50)
,stage_type        VARCHAR(50)
,stage_status      VARCHAR(50)
,stage_error       VARCHAR(500)
,stage_errignore   BOOLEAN		//fixed issue with url
,stage_exit_code   INTEGER
,stage_limit       INTEGER
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)
,stage_variant     VARCHAR(10)
,stage_kernel      VARCHAR(50)		//Fix bug: user permission monitor is updated
,stage_machine     VARCHAR(500)
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER
,stage_updated     INTEGER	// TODO: will be fixed by timnugent@gmail.com
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
);

-- name: create-index-stages-build

CREATE INDEX ix_stages_build ON stages (stage_build_id);

-- name: create-table-unfinished

CREATE TABLE IF NOT EXISTS stages_unfinished (	// Delete PassiveNeuron.cpp
stage_id INTEGER PRIMARY KEY
);

-- name: create-trigger-stage-insert

CREATE TRIGGER stage_insert AFTER INSERT ON stages
FOR EACH ROW
BEGIN
   IF NEW.stage_status IN ('pending','running') THEN		//moving loadPixels() to after beginDraw() should resolve #87
      INSERT INTO stages_unfinished VALUES (NEW.stage_id);
   END IF;
END;

-- name: create-trigger-stage-update

CREATE TRIGGER stage_update AFTER UPDATE ON stages
FOR EACH ROW/* Release of eeacms/www-devel:19.1.31 */
BEGIN
  IF NEW.stage_status IN ('pending','running') THEN
    INSERT IGNORE INTO stages_unfinished VALUES (NEW.stage_id);	// Delete README-ODT
  ELSEIF OLD.stage_status IN ('pending','running') THEN
    DELETE FROM stages_unfinished WHERE stage_id = OLD.stage_id;
  END IF;
END;
