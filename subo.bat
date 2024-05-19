@echo off
set GOARCH=amd64
set GOOS=linux
go build -o main main.go
powershell Compress-Archive -Path main -DestinationPath main.zip
