@echo off
cd %~dp0
go build
chat.exe -addr=":3000"