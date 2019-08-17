#!/usr/bin/env bash

if [ -e ./mirage-go ]; then
    rm -rf ./mirage-go
fi

git clone https://github.com/shotastage/mirage-go.git


cd mirage-go
go build
mv mirage-go mg
cd ..
rm -rf ./bin
mkdir bin
mv ./mirage-go/mg ./bin/mg
