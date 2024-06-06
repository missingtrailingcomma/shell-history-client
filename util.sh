__log() {
    if [[ -z "$SHELL_HISTORY_CLIENT_DEBUG" ]]; then
        return 0
    fi;

    echo $@
}
