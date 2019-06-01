#!/bin/bash

# generate govst.h
go build --buildmode=c-shared -o govst.so govst.go

# compile into shared library
go build -x --buildmode=c-shared -ldflags '-extldflags -Wl,-soname,govst.so' -o govst.so
