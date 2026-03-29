#!/usr/bin/env bash

dir=$(dirname "$0")
script=$(basename "$0")
cd "$dir"

gcc -o cfunction cfunction.c ccall.o

