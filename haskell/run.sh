#!/bin/bash

set -eu

if [[ "$#" -ne 3 ]]; then
  echo "Usage:"
  echo "  ./run.sh <year> <day> <part>"
  exit 1
fi

YEAR="${1}"
DAY="${2}"
PART="${3}"

runhaskell "${YEAR}-day-${DAY}-part-${PART}.hs" < "../input/${YEAR}/${DAY}.txt"
