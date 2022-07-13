#!/bin/bash
set -e

build() {
    go install
}

test::all() {
    go test -v ./...
}

tests() {
    go test -v -run TestIndent
}

"$@"
