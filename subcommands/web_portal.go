package subcommands

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"shell_history_client/defs"
	"time"

	"google.golang.org/protobuf/encoding/prototext"

	pb "shell_history_client/proto"
)

type pageData struct {
	Title string
	Rows  []row
}

type row struct {
	ExecutionStatus  int
	ExecutionTimeStr string
	CommandText      string

	Command *pb.Command
}

func WebPortal(env defs.EnvInfo) error {
	inputCacheFilePath := path.Join(env.User.HomeDir, defs.DOTFILE_FOLDER, defs.INPUT_CACHE_FILE)

	// Handle root URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := serveIndex(inputCacheFilePath, w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Start the HTTP server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	return nil
}

func serveIndex(inputCacheFilePath string, w http.ResponseWriter, r *http.Request) error {
	cmdList, err := readCommands(inputCacheFilePath)
	if err != nil {
		return fmt.Errorf("readCommands(%q): %v", inputCacheFilePath, err)
	}

	cmdLists, err := mergeCommands(cmdList)
	if err != nil {
		return fmt.Errorf("mergeCommands(): %v", err)
	}

	var rows []row
	for _, cmd := range cmdLists.Commands {
		convertedTime := cmd.ExecutionTime.AsTime()

		rows = append(rows, row{
			ExecutionStatus:  int(cmd.ExecutionStatus),
			ExecutionTimeStr: convertedTime.Format(time.RFC3339Nano),
			CommandText:      cmd.Text,
			Command:          cmd,
		})
	}

	tmpl, err := template.ParseFiles("web_portal/index.html")
	if err != nil {
		return fmt.Errorf(`template.ParseFiles("web_portal/index.html"): %v`, err)
	}

	// Data to pass to the template
	data := pageData{
		Title: "Gummy Bear 🍬🐻",
		Rows:  rows,
	}

	// Execute the template with the data
	err = tmpl.Execute(w, data)
	if err != nil {
		return fmt.Errorf("tmpl.Execute(): %v", err)
	}

	return nil
}

func readCommands(inputCacheFilePath string) (*pb.CommandList, error) {
	cmdList := pb.CommandList{}

	// Read existing commands.
	if fileExists(inputCacheFilePath) {
		inputCache, err := os.ReadFile(inputCacheFilePath)
		if err != nil {
			return nil, fmt.Errorf("io.ReadAll(%q): %v", inputCacheFilePath, err)
		}

		if string(inputCache) != "" {
			uo := prototext.UnmarshalOptions{
				DiscardUnknown: true,
			}

			if err := uo.Unmarshal(inputCache, &cmdList); err != nil {
				return nil, fmt.Errorf("proto.Unmarshal(): %v", err)
			}
		}
	}

	return &cmdList, nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func mergeCommands(cmdList *pb.CommandList) (*pb.CommandList, error) {
	cmds := cmdList.Commands

	var mergedCmds []*pb.Command

	cmdIds := make(map[string]bool)
	for _, cmd := range cmds {
		if _, ok := cmdIds[cmd.Id]; ok {
			for i := len(mergedCmds) - 1; i >= 0; i-- {
				if mergedCmds[i].Id == cmd.Id {
					mergedCmds[i].ExecutionStatus = cmd.ExecutionStatus
					break
				}
			}
		} else {
			mergedCmds = append(mergedCmds, cmd)
			cmdIds[cmd.Id] = true
		}
	}

	var reversed []*pb.Command
	for i := len(mergedCmds) - 1; i >= 0; i-- {
		reversed = append(reversed, mergedCmds[i])
	}

	return &pb.CommandList{Commands: reversed}, nil
}
