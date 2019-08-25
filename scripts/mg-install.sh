#!/usr/bin/env bash

function clean() {
    cd $HOME/.mirage-tool/
    rm -rf mirage-go
}

cd # Move home directory

if [ -e .mirage-tool/ ]; then
    echo "MIRAGE installation is already exists!"
    echo "Exit installer..."
    exit 1
fi


# Prepare directory
mkdir .mirage-tool/
cd .mirage-tool/

# Download
echo "⬇️  Downloading..."
git clone --quiet https://github.com/shotastage/mirage-go.git >> /dev/null

# Build 
cd mirage-go/
echo "🔄  Building..."
go build >> /dev/null

mv mirage-go mg
cd ..

mkdir bin
mv ./mirage-go/mg ./bin/mg 



mv ./mirage-go/scripts/mg-update.sh ./bin/
cd ./bin/
mv mg-update.sh mg-update


clean


if [ -f $HOME/.bash_profile ]; then
    echo "Writing Bash path..."
    echo 'export PATH="$HOME/.mirage-tool/bin:$PATH"' >> $HOME/.bash_profile
    echo "If you use other Shell, please set path manually."
fi
