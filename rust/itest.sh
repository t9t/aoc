#!/bin/bash

set -eu

cargo test test_all -- --ignored
