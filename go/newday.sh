#!/bin/bash

set -eu

if [[ $# -ne 1 ]]; then
  echo 'Usage: ./newday.sh <day>'
  exit 1
fi

DAY=${1}

cat "year2018/template.go" | sed "s/__daynum__/${DAY}/g" | sed "s/__remove__//g" | sed "s/5_5_2_1/${DAY}/g" > "year2018/day${DAY}.go"
cat "year2018/template_test.go" | sed "s/__daynum__/${DAY}/g" | sed "s/__remove__//g" | sed "s/5_5_2_1/${DAY}/g" > "year2018/day${DAY}_test.go"
