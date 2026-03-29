#!/usr/bin/env bash

dir=$(dirname "$0")
script=$(basename "$0")
cd "$dir"

rm -rf *.o
rm -rf *.a
rm -rf cfunction ccall.h
