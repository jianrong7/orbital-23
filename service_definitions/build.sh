#!/bin/bash
export GOOS=linux
export GOARCH=amd64

cd ./idlmanagement
go build .

cd ../service1v1
go build .

cd ../service1v2
go build .

cd ../service2v1
go build .