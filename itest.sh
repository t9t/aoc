#!/bin/sh

set -eu

go test -tags=itest -count=1 ./...
