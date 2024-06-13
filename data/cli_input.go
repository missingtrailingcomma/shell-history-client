package data

import (
	"os/user"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Input struct {
	CommandInput CommandInput
	ContextInfo  EnvInfo
}

type CommandInput struct {
	CommandId       string
	CommandText     string
	ExecutionStatus int
	WorkingDir      string
	ExecutionTime   *timestamppb.Timestamp
	Pid             int
	Ppid            int
}

type EnvInfo struct {
	User  *user.User
	Debug bool
}
