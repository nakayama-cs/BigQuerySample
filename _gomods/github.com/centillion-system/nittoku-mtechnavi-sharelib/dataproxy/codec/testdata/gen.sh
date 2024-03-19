#!/bin/bash

set -e -u

basedir="$(dirname "${0}")"
basedir="$(realpath -m "${basedir}")"

cd "${basedir}"

if ls *.go >/dev/null 2>&1; then
    rm *.go
fi

protoc \
    --proto_path='.' \
    --proto_path="${basedir}/../../../.local/opt/googleapis" \
    --proto_path="${basedir}/../../../.local/opt/protoc-gen-validate" \
    --go_opt=paths=source_relative \
    --go_out=. \
    --go-grpc_opt=paths=source_relative \
    --go-grpc_out=. \
    --validate_opt=lang=go,paths=source_relative \
    --validate_out=. \
    *.proto
