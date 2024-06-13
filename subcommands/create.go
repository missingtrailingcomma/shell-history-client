package subcommands

import (
	"fmt"
	"os"
	"path"
	"shell_history_client/data"
	"strings"

	"google.golang.org/protobuf/encoding/prototext"

	pb "shell_history_client/proto"
)

func Create(env data.EnvInfo, cmd *pb.Command) error {
	inputCacheFilePath := path.Join(env.User.HomeDir, data.DOTFILE_FOLDER, data.INPUT_CACHE_FILE)
	f, err := os.OpenFile(inputCacheFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("os.OpenFile(%q): %v", inputCacheFilePath, err)
	}
	defer f.Close()

	// Create or append, not overwrite so file I/O is limited.
	mo := &prototext.MarshalOptions{
		Indent: strings.Repeat(" ", 2),
	}
	pbtxt, err := mo.Marshal(&pb.CommandList{Commands: []*pb.Command{cmd}})
	if err != nil {
		return fmt.Errorf("prototext.MarshalOptions.Marshal(): %v", err)
	}
	if _, err := f.Write(pbtxt); err != nil {
		return fmt.Errorf("f.Write(toWrite): %v", err)
	}

	return nil
}
