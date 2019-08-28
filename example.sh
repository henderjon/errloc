#!/usr/bin/env sh

set -e
echo "docker run --rm -v $(PWD):/usr/src/myapp -w /usr/src/myapp golang:1.13rc1 go run example/main.go"
