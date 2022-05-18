#!/bin/bash

set -eu

if [ $# -lt 1 ]; then
  echo 'Usage: ./itest.sh <year>'
  exit 1
fi

YEAR="${1}"

go test -tags="itest_${YEAR}" -count=1 ./...
