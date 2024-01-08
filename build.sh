#!/usr/bin/env bash

SCRIPT_PATH="$(realpath "$0")"
PROJECT_DIR="$(dirname "${SCRIPT_PATH}")"
CMD_DIR="${PROJECT_DIR}/cmd"
BIN_DIR="${PROJECT_DIR}/bin"

mkdir -p "${BIN_DIR}"

cd "${CMD_DIR}"
go build

cd "${PROJECT_DIR}"
mv "${CMD_DIR}/cmd" "${BIN_DIR}/mtg-collector"


