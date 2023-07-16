#!/bin/bash
export GOOS=linux
export GOARCH=amd64

cd ./idlmanagement
rm idlmanagement
go build .

cd ../service1v1
rm service1v1
go build .

cd ../service1v2
rm service1v2
go build .

cd ../service2v1
rm service2v1
go build .