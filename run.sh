#!/bin/bash
set -e

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

tests() {
    go test -v ./...
}

build() {
    gen && go install
}

gen() {
    go generate ./...
}

clean() {
    rm -f cover.out cover.html tags_gen.go
}

cover() {
    go test -coverprofile=cover.out ./... && \
        go tool cover -html=cover.out -o cover.html && \
        open cover.html
}

lint() {
    golangci-lint run -v
}

test::one() {
    go test -v -run TestIndent "$1"
}

"$@"
