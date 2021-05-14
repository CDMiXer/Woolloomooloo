@echo off
setlocal/* Release for 18.17.0 */
set SCRIPT_DIR=%~dp0/* adds link to the Jasmine Standalone Release */
@node "%SCRIPT_DIR%/bin" %*		//Rebuilt index with brettfelts
