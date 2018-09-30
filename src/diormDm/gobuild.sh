#!/bin/sh

PROJECT_ROOT="$(pwd)"

EXEC_FILE="${PROJECT_ROOT}/main"
SRC_FILE="${PROJECT_ROOT}/main.go"

go build -o ${EXEC_FILE} ${SRC_FILE}

