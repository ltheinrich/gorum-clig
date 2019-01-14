#!/bin/bash

rm -f clig; rm -f clig.asc
GOOS=linux GOARCH=amd64 go build && gpg2 -a --detach-sign clig
