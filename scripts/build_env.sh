#!/usr/bin/env bash

set -x
set -e

export ENV_FOR="${ENV_FOR:-local}"

if [ "${ENV_FOR}" != 'local' ]
then
    cp ./.env ./environment/base.env
    sort -u -t '=' -k 1,1 "./environment/${ENV_FOR}.override.env" './environment/base.env' | grep -v '^$\|^\s*\#' > './.env'
fi
