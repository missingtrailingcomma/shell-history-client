# Shell History

## Install

TODO

## Uninstall

TODO

## Dev

```sh
mkdir -p $HOME/.yak1_gummybear
touch $HOME/.yak1_gummybear/cmd_cache.json

SHELL_HISTORY_CLIENT_DEBUG=1
```

```sh
bazel build //:shell-history-client && source shell-history-client.sh
```

### Debug Mode

```
SHELL_HISTORY_CLIENT_DEBUG=1
```

### Update Golang Deps

```sh
bazel run //:gazelle
```

## Dependencies

- https://github.com/rcaloras/bash-preexec

## Credit

- https://github.com/rcaloras/bashhub-client
- chatgpt.com
