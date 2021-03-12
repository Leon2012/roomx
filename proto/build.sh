#!/bin/sh

for x in **/*.proto; do protoc --go_out=plugins=grpc,paths=source_relative:. $x; done

#protoc --go_out=plugins=grpc,paths=source_relative:. ./proto/message.proto

#protoc --proto_path=common/ --go_out=plugins=grpc:common common.proto

#protoc --proto_path=common/ --proto_path=message/ --go_out=plugins=grpc:message message.proto