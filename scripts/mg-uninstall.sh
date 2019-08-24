#!/usr/bin/env bash


function uninstall() {
    # rm -rf .mirage-tool/
    echo "MIRAGE Go has been uninstalled!"
    echo "Please remove PATH manually."
}

function yesno {
    while true; do
        echo "Are you sure to uninstall MIRAGE Go?"
        echo "This operation can not be recovered."
        echo -n "$* Please answer yes or no [y/n]: "
        read ANS
        case $ANS in
        [Yy]*)
            uninstall
            ;;  
        [Nn]*)
            echo "Canceled"
            exit 1
            ;;
        *)
            echo "Please answer Y or N."
            ;;
    esac
  done
}

cd
yesno
