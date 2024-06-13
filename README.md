# Shell History

## Install

TODO

## Uninstall

TODO

## Dev

### Setup

```sh
mkdir -p $HOME/.yak1_gummybear

SHELL_HISTORY_CLIENT_DEBUG=1
```

### Source the latest version shell script

```sh
go build -o dist/gummybear main.go && source shell-history-client.sh
```

### Serve web portal

```sh
go run main.go --mode=web_portal
```

```sh
protoc --proto_path=./ --go_out=paths=source_relative:./ --go-grpc_out=paths=source_relative:./ $(find proto -name "*.proto") 
```

## Dependencies

- https://github.com/rcaloras/bash-preexec

## Credit

- https://github.com/rcaloras/bashhub-client
- chatgpt.com
