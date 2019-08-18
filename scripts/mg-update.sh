#!/usr/bin/env bash


function install_body() {
    mkdir .mirage-tool/
    cd .mirage-tool/
    echo "Downloading update..."
    git clone --quiet https://github.com/shotastage/mirage-go.git > /dev/null
    cd mirage-go/
    echo "Building..."
    go build 
    mv mirage-go mg
    cd ..

    mkdir bin
    mv ./mirage-go/mg ./bin/mg 

    mv ./mirage-go/scripts/mg-update.sh ./bin/
    cd ./bin/
    mv mg-update.sh mg-update
}



function restore_config() {
    rm -rf $HOME/.mirage-tool-tmp/
    echo "Restore config:: NOW UNDER CONSTRUCTION"
}


function restore_data() {
    echo "Restore data:: NOW UNDER CONSTRUCTION"
}

# HOME DIRECTORY
cd

# Backup files
mv .mirage-tool/ .mirage-tool-tmp/

rm -rf .mirage-tool/

install_body
restore_config
restore_data
