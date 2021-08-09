#!/bin/bash

set -eu

go test -tags=itest -count=1 ./...
