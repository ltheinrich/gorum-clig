#!/bin/bash

rm -rf bin
GOOS=linux GOARCH=amd64 go build -o bin/clig-linux-amd64
GOOS=freebsd GOARCH=amd64 go build -o bin/clig-freebsd-amd64
cd bin; gpg2 -a --detach-sign clig-linux-amd64 && gpg2 -a --detach-sign clig-freebsd-amd64
