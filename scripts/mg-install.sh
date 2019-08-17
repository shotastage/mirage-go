#!/usr/bin/env bash

cd # Move home directory

if [ -e .mirage-tool/ ]; then
    echo "MIRAGE installation is already exists!"
    echo "Exit installer..."
    exit 1
fi


mkdir .mirage-tool/
cd .mirage-tool/
git clone https://github.com/shotastage/mirage-go.git
cd mirage-go/
echo "Building..."
go build
mv mirage-go mg
cd ..

mkdir bin
mv ./mirage-go/mg ./bin/mg 




if [ -f $HOME/.bash_profile ]; then
    echo "Writing path..."
    echo 'export PATH="$HOME/.mirage-tool/bin:$PATH"' >> $HOME/.bash_profile
fi