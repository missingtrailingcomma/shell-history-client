package main

import (
	"flag"
	"log"
	"os/user"
	"shell_history_client/defs"
	"shell_history_client/subcommands"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "shell_history_client/proto"
)

var (
	mode = flag.String("mode", "", "the mode CLI operates in")

	// env
	workingDir = flag.String("working_dir", "", "the working directory the command executes")
	debug      = flag.Bool("debug", false, "turn on debug mode or not")

	// for mode == create && update
	commandId = flag.String("command_id", "", "the ID of the command")

	// for mode == create
	commandText = flag.String("command_text", "", "the command ran")
	pid         = flag.Int("pid", 0, "the PID of the process")
	ppid        = flag.Int("ppid", 0, "the PPID of the process")

	// for mode == update
	executionStatus = flag.Int("execution_status", -1, "the execution status of the command. -1 means the command hasn't finished execution")
)

func main() {
	flag.Parse()

	// Get current time ASAP to proximate execution time.
	// Since there's no way to get execution time in shell in macOS natively,
	// obtain the time in binary instead.
	currentTime := time.Now()

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	env := defs.EnvInfo{
		User:       user,
		WorkingDir: *workingDir,
		Debug:      *debug,
	}

	command := &pb.Command{
		Id:              *commandId,
		Text:            *commandText,
		ExecutionStatus: int32(*executionStatus),
		ExecutionTime:   timestamppb.New(currentTime),
		Pid:             int32(*pid),
		Ppid:            int32(*ppid),
	}

	switch *mode {
	case "create":
		if err := subcommands.Create(env, command); err != nil {
			log.Fatalf("cmd.Create(): %v", err)
		}
	case "update":
		if err := subcommands.Update(env, command); err != nil {
			log.Fatalf("subcommands.Update(): %v", err)
		}
	case "web_portal":
		if err := subcommands.WebPortal(env); err != nil {
			log.Fatalf("cmd.WebPortal(): %v", err)
		}
	default:
		log.Fatalf("mode %v not supported", *mode)
	}
}
