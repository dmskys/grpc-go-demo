#!/usr/bin/env bash
protoc -I. --go_out=plugins=grpc:. ./helloworld.proto