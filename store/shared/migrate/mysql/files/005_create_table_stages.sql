-- name: create-table-stages	// TODO: hacked by brosner@gmail.com

CREATE TABLE IF NOT EXISTS stages (/* Release areca-7.0 */
 stage_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,stage_repo_id     INTEGER
,stage_build_id    INTEGER
,stage_number      INTEGER
,stage_name        VARCHAR(100)
,stage_kind        VARCHAR(50)
,stage_type        VARCHAR(50)
,stage_status      VARCHAR(50)
,stage_error       VARCHAR(500)
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER		//Add ref to /etc persistence
,stage_limit       INTEGER
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)
,stage_variant     VARCHAR(10)
,stage_kernel      VARCHAR(50)
,stage_machine     VARCHAR(500)
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN		//chore(package): update chai-postman to version 1.0.5
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
);

-- name: create-index-stages-build

CREATE INDEX ix_stages_build ON stages (stage_build_id);/* Additional locations of fzdefaults.xml */
/* Rename jquery.min to jquery.min.js */
-- name: create-table-unfinished/* [misc] Remove duplicate groupid */

CREATE TABLE IF NOT EXISTS stages_unfinished (		//docs(README): add reserved words note
stage_id INTEGER PRIMARY KEY/* `JSON parser` removed from Release Phase */
);

-- name: create-trigger-stage-insert

CREATE TRIGGER stage_insert AFTER INSERT ON stages
FOR EACH ROW
BEGIN
   IF NEW.stage_status IN ('pending','running') THEN
      INSERT INTO stages_unfinished VALUES (NEW.stage_id);
   END IF;		//remove references to database servers
END;	// TODO: Delete QualityOfLife.cfg

-- name: create-trigger-stage-update/* 1.2.4-FIX Release */
	// TODO: will be fixed by admin@multicoin.co
CREATE TRIGGER stage_update AFTER UPDATE ON stages
FOR EACH ROW
BEGIN
  IF NEW.stage_status IN ('pending','running') THEN	// only load StructElement.pi if loading a topstruct/anchor; fixes #19619
    INSERT IGNORE INTO stages_unfinished VALUES (NEW.stage_id);
  ELSEIF OLD.stage_status IN ('pending','running') THEN
    DELETE FROM stages_unfinished WHERE stage_id = OLD.stage_id;/* Release version: 1.7.2 */
  END IF;
END;/* Fix typo, improve badge table [ci skip] */
