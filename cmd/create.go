package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"shell_history_client/data"

	"google.golang.org/protobuf/encoding/prototext"

	pb "shell_history_client/proto"
)

func Create(env data.EnvInfo, cmd *pb.Command) error {
	inputCacheFilePath := path.Join(env.User.HomeDir, data.DOTFILE_FOLDER, data.INPUT_CACHE_FILE)

	// cmdList := pb.CommandList{}

	// // Read existing commands.
	// if fileExists(inputCacheFilePath) {
	// 	inputCache, err := os.ReadFile(inputCacheFilePath)
	// 	if err != nil {
	// 		log.Fatalf("io.ReadAll(%q): %v", inputCacheFilePath, err)
	// 	}

	// 	if string(inputCache) != "" {
	// 		if err := proto.Unmarshal(inputCache, &cmdList); err != nil {
	// 			log.Fatalf("proto.Unmarshal(): %v", err)
	// 		}
	// 	}
	// }

	// // Add the latest command.
	// cmdList.Commands = append(cmdList.Commands, cmd)

	// Add to file.
	f, err := os.OpenFile(inputCacheFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("os.OpenFile(%q): %v", inputCacheFilePath, err)
	}
	defer f.Close()

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
