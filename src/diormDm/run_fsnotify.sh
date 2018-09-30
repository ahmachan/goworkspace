#!/bin/sh

CMD_ROOT="/opt/code"
#PROJECT_ROOT="${CMD_ROOT}/martini-test3"
PROJECT_ROOT="${CMD_ROOT}/diorm-demo"

${CMD_ROOT}/xmage_fsnotify -f ${PROJECT_ROOT}/main.go -o ${PROJECT_ROOT}/main

