#!/bin/bash

source "dependencies/bash-preexec/bash-preexec.sh"
source "util.sh"



__preexec_hook() {
    local _PID=$$
    local _PPID=$PPID
    local _DEBUG="false"
    
    # Used for track command execution across __preexec_hook and __precmd_hook.
    __command_id=$(uuidgen)

    if [ -n "$SHELL_HISTORY_CLIENT_DEBUG" ]; then
        _DEBUG="true"
    fi
    
    # TODO: can we use `RUNFILES` here?
    bazel-bin/shell_history_client_/shell_history_client --mode=create \
        -command_id=$__command_id \
        -command_text=$1 \
        -working_dir=$PWD \
        -pid=$_PID \
        -ppid=$_PPID \
        -debug=$_DEBUG
}

__precmd_hook() {
    local __exit_status="$?"    
    unset __command_id
}

# Clear precmd_functions in debug mode so only precmd_functions registered in this package are in effect.
if [ -n "$SHELL_HISTORY_CLIENT_DEBUG" ]; then
    precmd_functions=()
    preexec_functions=()
fi
preexec_functions+=(__preexec_hook)
precmd_functions+=(__precmd_hook)
