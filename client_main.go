package main

import (
	"flag"
	"fmt"
	"log"
	"os/user"
	"time"

	"shell_history_client/cmd"
	"shell_history_client/data"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	mode = flag.String("mode", "", "the mode CLI operates in")

	commandId       = flag.String("command_id", "", "the ID of the command")
	commandText     = flag.String("command_text", "", "the command ran")
	executionStatus = flag.Int("execution_status", 0, "the execution status of the command")
	workingDir      = flag.String("working_dir", "", "the working directory the command executes")
	pid             = flag.Int("pid", 0, "the PID of the process")
	ppid            = flag.Int("ppid", 0, "the PPID of the process")

	debug = flag.Bool("debug", false, "turn on debug mode or not")
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

	input := data.Input{
		CommandInput: data.CommandInput{
			CommandId:       *commandId,
			CommandText:     *commandText,
			ExecutionStatus: *executionStatus,
			WorkingDir:      *workingDir,
			ExecutionTime:   timestamppb.New(currentTime),
			Pid:             *pid,
			Ppid:            *ppid,
		},
		ContextInfo: data.EnvInfo{
			User:  user,
			Debug: *debug,
		},
	}

	switch *mode {
	case "create":
		if err := cmd.Create(input); err != nil {
			log.Fatalf("cmd.Create(): %v", err)
		}
	case "web_portal":
		if err := cmd.WebPortal(input); err != nil {
			log.Fatalf("cmd.WebPortal(): %v", err)
		}
	default:
		log.Fatalf("mode %v not supported", *mode)
	}

	fmt.Printf("input: %+v\n\n", input)
}
