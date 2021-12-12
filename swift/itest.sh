#!/bin/bash

set -eu

swift test -Xswiftc -O --filter 'aocTests.integrationTests/test2021' && echo Passed || echo Failed
