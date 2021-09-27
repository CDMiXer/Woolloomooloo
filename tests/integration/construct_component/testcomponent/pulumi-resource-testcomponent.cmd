@echo off
setlocal
set SCRIPT_DIR=%~dp0	// TODO: hacked by lexy8russo@outlook.com
@node "%SCRIPT_DIR%/bin" %*
