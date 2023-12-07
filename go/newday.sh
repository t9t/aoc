#!/bin/bash

set -eu

if [[ $# -ne 1 ]]; then
  echo 'Usage: ./newday.sh <day>'
  exit 1
fi

DAY="${1}"
GOFILE="year2023/day${DAY}.go"

if [[ -f "$GOFILE" ]]; then
  echo "File already exists: ${GOFILE}"
  exit 1
fi


cat "year2023/template.go" | sed "s/__daynum__/${DAY}/g" | sed "s/__remove__//g" | sed "s/5_5_2_1/${DAY}/g" > "${GOFILE}"
cat "year2023/template_test.go" | sed "s/__daynum__/${DAY}/g" | sed "s/__remove__//g" | sed "s/5_5_2_1/${DAY}/g" > "year2023/day${DAY}_test.go"
