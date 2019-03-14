#!/bin/sh

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w -s' -mod=vendor -o build/deploy_linux_amd64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -ldflags '-w -s' -mod=vendor -o build/deploy_windows_amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -ldflags '-w -s' -mod=vendor -o build/deploy_mac_amd64
