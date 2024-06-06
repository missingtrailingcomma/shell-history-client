source "dependencies/bash-preexec/bash-preexec.sh"
source "util.sh"

__precmd_hook() {
    local _PID=$$
    local _PPID=$PPID
    local _DEBUG="false"

    if [ -n "$SHELL_HISTORY_CLIENT_DEBUG" ]; then
        _DEBUG="true"
    fi
    
    bazel-bin/shell-history-client_/shell-history-client \
        -pid=$_PID \
        -ppid=$_PPID \
        -debug=$_DEBUG
}

# Clear precmd_functions in debug mode so only precmd_functions registered in this package are in effect.
if [ -n "$SHELL_HISTORY_CLIENT_DEBUG" ]; then
    precmd_functions=()
fi
precmd_functions+=(__precmd_hook)
