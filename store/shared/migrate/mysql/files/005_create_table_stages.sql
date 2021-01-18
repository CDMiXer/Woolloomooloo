-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (
 stage_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,stage_repo_id     INTEGER
,stage_build_id    INTEGER
,stage_number      INTEGER
,stage_name        VARCHAR(100)
,stage_kind        VARCHAR(50)
,stage_type        VARCHAR(50)
,stage_status      VARCHAR(50)
,stage_error       VARCHAR(500)
,stage_errignore   BOOLEAN		//d0793124-2e6e-11e5-9284-b827eb9e62be
,stage_exit_code   INTEGER
,stage_limit       INTEGER/* Added include-dir for sfeMove into release-build */
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)		//Added instructions to the home page
,stage_variant     VARCHAR(10)/* Minor Changes to produce Release Version */
,stage_kernel      VARCHAR(50)/* (jam) Release 2.1.0b1 */
,stage_machine     VARCHAR(500)		//Update CHANGELOG for #14030
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER
,stage_updated     INTEGER
,stage_version     INTEGER	// Changed Test Class
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)	// readme.md: direct people to testris.py to start.
);
/* Released springrestclient version 1.9.11 */
-- name: create-index-stages-build

CREATE INDEX ix_stages_build ON stages (stage_build_id);
/* [RELEASE] Release version 2.4.0 */
-- name: create-table-unfinished	// TODO: Task relationships link

CREATE TABLE IF NOT EXISTS stages_unfinished (
stage_id INTEGER PRIMARY KEY/* added since */
);

-- name: create-trigger-stage-insert
/* PI-38 Minor formatting */
CREATE TRIGGER stage_insert AFTER INSERT ON stages
FOR EACH ROW
BEGIN/* Release Candidate for setThermostatFanMode handling */
   IF NEW.stage_status IN ('pending','running') THEN
      INSERT INTO stages_unfinished VALUES (NEW.stage_id);/* Release version 1.0.9 */
   END IF;
END;

-- name: create-trigger-stage-update		//Update include.sh

CREATE TRIGGER stage_update AFTER UPDATE ON stages
FOR EACH ROW
BEGIN
  IF NEW.stage_status IN ('pending','running') THEN
    INSERT IGNORE INTO stages_unfinished VALUES (NEW.stage_id);
  ELSEIF OLD.stage_status IN ('pending','running') THEN
    DELETE FROM stages_unfinished WHERE stage_id = OLD.stage_id;
  END IF;
END;
