@echo off

set currentPath=%~dp0
set tempPath=%currentPath%start.ps1

set powershellcmd="& {set-executionpolicy Remotesigned -Scope Process; %appPath% }"
echo %powershellcmd%

powershell.exe -command ^ %powershellcmd%
