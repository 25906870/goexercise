@echo off
set GOPATH=%cd%
cd ./bin/windows
@echo on
go build ../../src/main.go
cd ..
pause