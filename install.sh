#!/bin/bash -e

set -e

install() {
    setup_files
}

setup_files() {
    mkdir -p $HOME/.yak1_gummybear
    touch $HOME/.yak1_gummybear/cmd_cache.json
    cd $HOME/.yak1_gummybear
}
