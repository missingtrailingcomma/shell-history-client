package subcommands

import (
	"shell_history_client/defs"

	pb "shell_history_client/proto"
)

func Update(env defs.EnvInfo, cmd *pb.Command) error {
	return Create(env, cmd)
}
