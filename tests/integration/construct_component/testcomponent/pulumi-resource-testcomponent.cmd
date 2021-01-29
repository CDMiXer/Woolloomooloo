@echo off/* Release 3.0.0 */
setlocal		//Merge "x-newest cleanup code with test. Fixes bug 1037337"
set SCRIPT_DIR=%~dp0		//186ff6fe-2e68-11e5-9284-b827eb9e62be
@node "%SCRIPT_DIR%/bin" %*
