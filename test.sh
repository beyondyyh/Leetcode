#/bin/bash
pwd
set -e
rm -rf coverage.txt

for dir in `go list ./algorithms/...`; do
    # echo $dir
    go test -coverprofile=profile.out -covermode=atomic $dir
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm -f profile.out
    fi
done

