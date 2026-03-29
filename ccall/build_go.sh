#!/usr/bin/env bash

dir=$(dirname "$0")
script=$(basename "$0")
cd "$dir"
# build go pkg to C
go build -o ccall.o -buildmode=c-archive exportfunc.go
