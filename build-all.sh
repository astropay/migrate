#!/bin/sh
cd cmd

rm -fr *.zip

# linux 64 bits
echo "building linux"
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o migrate
zip migrate-linux-amd64.zip migrate

# linux 32 bits
GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o migrate
zip migrate-linux-386.zip migrate

# macOS
echo "building macOS"
GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o migrate
zip migrate-darwin-64.zip migrate

# windows 64 bits
echo "building windows"
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o migrate
zip migrate-windows-64.zip migrate

# windows 32 bits
GOOS=windows GOARCH=386 go build -ldflags "-s -w" -o migrate
zip migrate-windows-32.zip migrate