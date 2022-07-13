#!/bin/bash
set -e

build() {
    go install
}

tests() {
    go test -v ./...
}

test::one() {
    go test -v -run TestIndent
}

"$@"
