#!/bin/bash

set -eu

echo "'session' cookie value: "
read -s SESH

INPUTDIR="./input"

mkdir -p "${INPUTDIR}"
for year in {2015..2020}; do
  for day in {1..25}; do
    echo "Fetching year ${year}; day ${day}"
    curl -f -o "${INPUTDIR}/${year}-${day}.txt" -H"Cookie: session=${SESH}" "https://adventofcode.com/${year}/day/${day}/input" 
  done
done
