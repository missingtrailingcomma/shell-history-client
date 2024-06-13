package cmd

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"shell_history_client/data"
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

func WebPortal(env data.EnvInfo) error {
	inputCacheFilePath := path.Join(env.User.HomeDir, data.DOTFILE_FOLDER, data.INPUT_CACHE_FILE)

	cmdList := pb.CommandList{}

	// Read existing commands.
	if fileExists(inputCacheFilePath) {
		inputCache, err := os.ReadFile(inputCacheFilePath)
		if err != nil {
			return fmt.Errorf("io.ReadAll(%q): %v", inputCacheFilePath, err)
		}

		if string(inputCache) != "" {
			uo := prototext.UnmarshalOptions{
				DiscardUnknown: true,
			}

			if err := uo.Unmarshal(inputCache, &cmdList); err != nil {
				return fmt.Errorf("proto.Unmarshal(): %v", err)
			}
		}
	}

	var rows []row
	for _, cmd := range cmdList.Commands {
		convertedTime := cmd.ExecutionTime.AsTime()

		rows = append(rows, row{
			ExecutionStatus:  int(cmd.ExecutionStatus),
			ExecutionTimeStr: convertedTime.Format(time.RFC3339Nano),
			CommandText:      cmd.Text,
			Command:          cmd,
		})
	}

	// Handle root URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("web_portal/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Data to pass to the template
		data := pageData{
			Title: "Gummy Bear 🍬🐻",
			Rows:  rows,
		}

		// Execute the template with the data
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Start the HTTP server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
