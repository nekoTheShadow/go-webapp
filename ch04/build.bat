@echo off

REM 環境変数BHT_APIKEYに次のURLで発行したAPIKEYを設定すること: https://words.bighugelabs.com/
REM 設定例: set BHT_APIKEY=1234567890abcd

set workspace=%~dp0

for %%f in (coolify domainnify sprinkle synonyms available) do (
    cd %workspace%
    del /Q %%f.exe
    cd ..\%%f
    go build -o %workspace%%%f.exe
)

cd %workspace%
synonyms.exe | sprinkle.exe | coolify.exe | domainnify.exe -t com -t net | available.exe
