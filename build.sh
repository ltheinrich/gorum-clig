#!/bin/bash

# linux amd64
GOOS=linux GOARCH=amd64 go build -o bin/clig-linux-amd64

# freebsd amd64
GOOS=freebsd GOARCH=amd64 go build -o bin/clig-freebsd-amd64
