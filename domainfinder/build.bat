@echo off

REM 環境変数BHT_APIKEYに次のURLで発行したAPIKEYを設定すること: https://words.bighugelabs.com/
REM 設定例: set BHT_APIKEY=1234567890abcd

set workspace=%~dp0

for %%f in (coolify domainnify sprinkle synonyms available) do (
    cd %workspace%
    del /Q lib\%%f.exe
    cd ..\%%f
    go build -o %workspace%lib\%%f.exe
)

cd %workspace%
go build
domainfinder.exe


REM cd %workspace%
REM synonyms.exe | sprinkle.exe | coolify.exe | domainnify.exe -t com -t net | available.exe
