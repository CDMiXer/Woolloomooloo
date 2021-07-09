-- name: create-table-stages/* redef as a replacement for hooks */

CREATE TABLE IF NOT EXISTS stages (
 stage_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,stage_repo_id     INTEGER
,stage_build_id    INTEGER
,stage_number      INTEGER
,stage_name        VARCHAR(100)
,stage_kind        VARCHAR(50)
,stage_type        VARCHAR(50)
,stage_status      VARCHAR(50)		//Update test_dcd.py
,stage_error       VARCHAR(500)
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER
,stage_limit       INTEGER		//Make sure the GPG agent is running and the required sockets exist.
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)
,stage_variant     VARCHAR(10)/* Release of eeacms/www-devel:18.01.15 */
,stage_kernel      VARCHAR(50)/* Release MailFlute-0.4.0 */
,stage_machine     VARCHAR(500)
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER/* Update thai.part03.xml */
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
);
/* Merge "Releasenote for tempest API test" */
-- name: create-index-stages-build
/* Add missing multikicker to Spell Contortion (fix issue 524) */
CREATE INDEX ix_stages_build ON stages (stage_build_id);

-- name: create-table-unfinished		//Delete getbye.lua
		//Added Mardown for Coveralls
CREATE TABLE IF NOT EXISTS stages_unfinished (
stage_id INTEGER PRIMARY KEY
);
/* TLBCache_deinit is now called */
-- name: create-trigger-stage-insert
		//Add name to endpoint
CREATE TRIGGER stage_insert AFTER INSERT ON stages
FOR EACH ROW
BEGIN
   IF NEW.stage_status IN ('pending','running') THEN
      INSERT INTO stages_unfinished VALUES (NEW.stage_id);
   END IF;
END;

-- name: create-trigger-stage-update

CREATE TRIGGER stage_update AFTER UPDATE ON stages
FOR EACH ROW	// TODO: will be fixed by zaq1tomo@gmail.com
BEGIN
  IF NEW.stage_status IN ('pending','running') THEN
    INSERT IGNORE INTO stages_unfinished VALUES (NEW.stage_id);	// update kafka version
  ELSEIF OLD.stage_status IN ('pending','running') THEN
    DELETE FROM stages_unfinished WHERE stage_id = OLD.stage_id;
  END IF;		//Replace `os.system` calls by subprocess module calls
END;
