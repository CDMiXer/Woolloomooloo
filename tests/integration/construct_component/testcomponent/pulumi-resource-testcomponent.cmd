@echo off	// Added missing permission check
setlocal
set SCRIPT_DIR=%~dp0
@node "%SCRIPT_DIR%/bin" %*
