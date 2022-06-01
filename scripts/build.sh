#!/usr/bin/env bash

set -euo pipefail
set -x
IFS=$'\n\t'

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOGC=off go build \
	-ldflags='-w -s -extldflags "-static"' -a -o ./.cache/myapp/myapp ./cmd/myapp/.
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 GOGC=off go build \
	-ldflags='-w -s -extldflags "-static"' -a -o ./.cache/myapp/myapp.exe ./cmd/myapp/.
