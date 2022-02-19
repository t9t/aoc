#!/bin/bash

set -eu

if [[ $# -ne 1 ]]; then
  echo 'Usage: ./newday.sh <day>'
  exit 1
fi

DAY=${1}

SOURCE_DIR="src"
SOURCE_FILE="${SOURCE_DIR}/day${DAY}.rs"
MAIN_FILE="${SOURCE_DIR}/main.rs"

cp -n "${SOURCE_DIR}/template.rs" "${SOURCE_FILE}"

sed -i "s/\/\*mod newday\*\//mod day${DAY};\n\/\*mod newday\*\//" "${MAIN_FILE}"
sed -i "s/ \/\*newday\*\/]/, day${DAY}::part1, day${DAY}::part2 \/\*newday\*\/]/" "${MAIN_FILE}"

git add "${SOURCE_FILE}" "${MAIN_FILE}"
