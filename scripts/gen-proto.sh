#!/bin/bash
CURRENT_DIR=$1
for x in $(find ${CURRENT_DIR}/template_protos/* -type d); do
  protoc -I=${x} -I=${CURRENT_DIR}/template_protos -I /usr/local/include --go_out=plugins=grpc:${CURRENT_DIR} ${x}/*.proto
done
