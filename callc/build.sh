#!/usr/bin/env bash

gcc -c callC.c -o callC.o
ar rcs libcallC.a callC.o
