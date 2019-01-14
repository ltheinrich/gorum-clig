#!/bin/bash

rm -rf bin; GOOS=linux GOARCH=amd64 go build -o bin/clig
cd bin; gpg2 -a --detach-sign clig
