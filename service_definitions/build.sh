
#!/bin/bash

export GOOS=linux
export GOARCH=amd64

cd ./idlmanagement
go build .

cd ../service1v2
go build .