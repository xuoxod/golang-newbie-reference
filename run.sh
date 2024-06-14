#!/usr/bin/env bash

pauseTime=1
sleepTime=2

clear
ls
sleep $pauseTime
rm -rf ./ahp
sleep $pauseTime
ls
sleep $pauseTime
go build -o ahp ./cmd/*.go
go mod tidy
sleep $pauseTime
ls
sleep $sleepTime
./ahp
