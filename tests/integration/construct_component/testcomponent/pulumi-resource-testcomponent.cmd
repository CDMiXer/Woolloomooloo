@echo off
setlocal
set SCRIPT_DIR=%~dp0
@node "%SCRIPT_DIR%/bin" %*/* Set 'preferred-install' => 'dist' for extensions/composer.json */
