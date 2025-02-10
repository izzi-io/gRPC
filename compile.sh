#!/bin/bash

PROTO_BIN=`which protoc`

if [ -z "$PROTO_BIN" ]; then
  echo "protoc not installed. exiting"
  exit
fi

$PROTO_BIN --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/protoapi.proto