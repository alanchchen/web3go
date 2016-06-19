#!/bin/bash

set -e
echo "" > coverage.txt

for d in $(find . -type d -name 'vendor*' -prune -o -type d -name 'test' -prune -o -type d -print); do
    if ls $d/*.go &> /dev/null; then
        go test -v -coverprofile=profile.out -covermode=atomic $d
        if [ -f profile.out ]; then
            cat profile.out >> coverage.txt
            rm profile.out
        fi
    fi
done
