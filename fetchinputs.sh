#!/bin/bash

set -eu

echo "'session' cookie value: "
read -s SESH

INPUTDIR="./input"

fetch() {
  year="${1}"
  day="${2}"
  YEARDIR="${INPUTDIR}/${year}"
  mkdir -p "${YEARDIR}"
  echo "Fetching year ${year}; day ${day}"
  curl -A 'blah blah blah' -f -o "${YEARDIR}/${day}.txt" -H"Cookie: session=${SESH}" "https://adventofcode.com/${year}/day/${day}/input" 
}

if [[ $# -eq 2 ]]; then
  fetch "${1}" "${2}"
  exit 0
fi

if [[ $# -ne 0 ]]; then
  echo 'Usage: ./fetchinputs.sh [year day]; either specify both or neither'
  exit 1
fi

for year in {2015..2022}; do
  for day in {1..25}; do
    fetch "${year}" "${day}"
  done
done
