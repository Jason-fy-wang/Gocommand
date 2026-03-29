#!/usr/bin/env bash

# clean up pkg
dir=$(dirname "$0")
script=$(basename "$0")
cd "$dir"
rm -rf *.o
rm -rf *.a
