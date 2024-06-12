package data

import "os/user"

type Input struct {
	CommandInput CommandInput
	ContextInfo  EnvInfo
}

type CommandInput struct {
	CommandId       string
	CommandText     string
	ExecutionStatus int
	WorkingDir      string
	ExecutionTime   string // TODO change to timestamp
	Pid             int
	Ppid            int
}

type EnvInfo struct {
	User  *user.User
	Debug bool
}
