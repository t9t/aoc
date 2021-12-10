#!/bin/bash

set -eu

if [[ $# -ne 1 ]]; then
  echo 'Usage: ./newday.sh <day>'
  exit 1
fi

DAY=${1}

SOURCE_DIR="Sources/aoc"
SOURCE_FILE="${SOURCE_DIR}/day${DAY}.swift"
TESTS_DIR="Tests/aocTests"
TESTS_FILE="${TESTS_DIR}/day${DAY}Tests.swift"
MAIN_FILE="${SOURCE_DIR}/main.swift"

cp -n "${SOURCE_DIR}/template.swift" "${SOURCE_FILE}"
gsed -i "s/TemplateDay/Day${DAY}/" "${SOURCE_FILE}"

cp -n "${TESTS_DIR}/templateTests.swift" "${TESTS_FILE}"
gsed -i "s/TemplateDay/Day${DAY}/" "${TESTS_FILE}"
gsed -i "s/TemplateDay/Day${DAY}/" "${TESTS_FILE}"
gsed -i "s/templateTests/day${DAY}Tests/" "${TESTS_FILE}"

gsed -i "s/\/\*newday\*\/]/        ${DAY}: Day${DAY}.init,\n\/\*newday\*\/]/" "${MAIN_FILE}"


git add "${SOURCE_FILE}" "${TESTS_FILE}" "${MAIN_FILE}"