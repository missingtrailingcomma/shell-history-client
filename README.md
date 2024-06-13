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

```sh
protoc --proto_path=./ --go_out=paths=source_relative:./ --go-grpc_out=paths=source_relative:./ $(find proto -name "*.proto") 
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
