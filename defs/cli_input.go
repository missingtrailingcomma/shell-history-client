package defs

import (
	"os/user"
)

type EnvInfo struct {
	User       *user.User
	WorkingDir string
	Debug      bool
}
