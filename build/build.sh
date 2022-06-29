#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail


if [ -z "${BIN}" ]; then
    echo "BIN must be set"
    exit 1
fi

if [ -z "${GOOS}" ]; then
    echo "GOOS must be set"
    exit 1
fi

if [ -z "${GOARCH}" ]; then
    echo "GOARCH must be set"
    exit 1
fi

if [[ -z "${OUTPUT_DIR:-}" ]]; then
  OUTPUT_DIR=.
fi
OUTPUT=${OUTPUT_DIR}/${BIN}

if [[ "${GOOS}" = "windows" ]]; then
  OUTPUT="${OUTPUT}.exe"
fi

go build \
    -o ${OUTPUT} \
    -installsuffix "static" \
    ./cmd/${BIN}