package main

import (
	"flag"
	"log"
	"os/user"
	"shell_history_client/data"
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

	// for mode == create
	commandId       = flag.String("command_id", "", "the ID of the command")
	commandText     = flag.String("command_text", "", "the command ran")
	executionStatus = flag.Int("execution_status", 0, "the execution status of the command")
	pid             = flag.Int("pid", 0, "the PID of the process")
	ppid            = flag.Int("ppid", 0, "the PPID of the process")
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

	env := data.EnvInfo{
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
	case "web_portal":
		if err := subcommands.WebPortal(env); err != nil {
			log.Fatalf("cmd.WebPortal(): %v", err)
		}
	default:
		log.Fatalf("mode %v not supported", *mode)
	}
}
