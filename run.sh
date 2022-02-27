#!/usr/bin/env bash

set -eo pipefail

DC="${DC:-exec}"

# If we're running in CI we need to disable TTY allocation for docker-compose
# commands that enable it by default, such as exec and run.
TTY=""
if [[ ! -t 1 ]]; then
    TTY="-T"
fi

# -----------------------------------------------------------------------------
# Helper functions start with _ and aren't listed in this script's help menu.
# -----------------------------------------------------------------------------

function _dc {
    export DOCKER_BUILDKIT=1
    docker-compose ${TTY} "${@}"
}

function _use_env {
    set -o allexport; . .env; set +o allexport
}

# ----------------------------------------------------------------------------

up() {
    reflex -c reflex.conf --decoration=fancy
}

up:myapp() {
    _use_env
    go run ./cmd/myapp "${@}"
}

build:myapp() {
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOGC=off go build \
    -ldflags='-w -s -extldflags "-static"' -a -o ./.cache/myapp/myapp ./cmd/myapp/.
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 GOGC=off go build \
    -ldflags='-w -s -extldflags "-static"' -a -o ./.cache/myapp/myapp.exe ./cmd/myapp/.
}

wire:myapp() {
    go install github.com/google/wire/cmd/wire@latest
    cd ./cmd/myapp && wire
}

up:compose() {
    docker-compose --profile main build --parallel
    docker-compose --profile main up
}

lint() {
    golangci-lint run --enable-all  --disable=wsl,varnamelen,testpackage,gomnd,exhaustivestruct
}

test() {
    go test -v -race -cover ./...
}

env:prod() {
    ENV_FOR=prod bash ./scripts/build_env.sh
}

# -----------------------------------------------------------------------------

function help {
    printf "%s <task> [args]\n\nTasks:\n" "${0}"

    compgen -A function | grep -v "^_" | cat -n
}

TIMEFORMAT=$'\nTask completed in %3lR'
time "${@:-help}"
