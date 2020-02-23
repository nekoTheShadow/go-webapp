@echo off

set WORK=%~dp0

cd cmds\backup
go build -o backup.exe
move backup.exe %WORK%
cd %WORK%

cd cmds\backupd
go build -o backupd.exe
move backupd.exe %WORK%
cd %WORK%
