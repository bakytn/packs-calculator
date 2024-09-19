#!/usr/bin/env bash
set -eo pipefail

script_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

cd "$script_dir"

docker build -f Dockerfile-base -t repartners/golang-base .
docker build -f Dockerfile-dev -t repartners/golang-dev .

cd -