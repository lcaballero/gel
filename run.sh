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

cover() {
    go test -coverprofile=cover.out ./... && \
        go tool cover -html=cover.out -o cover.html && \
        open cover.html
}

"$@"
