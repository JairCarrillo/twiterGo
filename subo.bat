@echo off
set GOARCH=amd64
set GOOS=linux
go build main.go
del main.zip
powershell Compress-Archive -Path main -DestinationPath main.zip
