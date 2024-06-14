#!/usr/bin/env bash

go build -o ahp ./cmd/*.go
go mod tidy

clear
