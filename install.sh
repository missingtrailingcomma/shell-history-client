#!/bin/bash -e

set -e

install() {
    setup_files
}

setup_files() {
    mkdir -p $HOME/.yak1_gummybear
    cd $HOME/.yak1_gummybear
}
