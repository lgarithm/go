#!/bin/sh
set -e

cd $(go env GOPATH)/pkg
[ -f mod.tar ] && rm mod.tar
tar -cf mod.tar mod
bzip2 mod.tar
