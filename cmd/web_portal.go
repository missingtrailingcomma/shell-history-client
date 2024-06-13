package cmd

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
	"shell_history_client/data"
	"text/template"
	"time"

	"github.com/bazelbuild/rules_go/go/tools/bazel"
)

type row struct {
	ExecutionStatus  int
	ExecutionTimeStr string
	CommandText      string

	Raw data.Input
}

type pageData struct {
	Title string
	Rows  []row
}

// TODO: change to take in context
func WebPortal(input data.Input) error {
	inputCacheFilePath := path.Join(input.ContextInfo.User.HomeDir, data.DOTFILE_FOLDER, data.INPUT_CACHE_FILE)

	var inputs []data.Input
	if fileExists(inputCacheFilePath) {
		inputCache, err := os.ReadFile(inputCacheFilePath)
		if err != nil {
			log.Fatalf("io.ReadAll(%q): %v", inputCacheFilePath, err)
		}

		if string(inputCache) != "" {
			if err := json.Unmarshal(inputCache, &inputs); err != nil {
				log.Fatalf("json.Unmarshal(): %v", err)
			}
		}
	}

	var rows []row
	for _, input := range inputs {
		convertedTime := input.CommandInput.ExecutionTime.AsTime()

		rows = append(rows, row{
			ExecutionStatus:  input.CommandInput.ExecutionStatus,
			ExecutionTimeStr: convertedTime.Format(time.RFC3339Nano),
			CommandText:      input.CommandInput.CommandText,
			Raw:              input,
		})
	}

	// Handle root URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Parse the template file

		runfilesPath, err := bazel.RunfilesPath()
		if err != nil {
			log.Fatalf("Could not locate data file: %v", err)
		}

		tmpl, err := template.ParseFiles(path.Join(runfilesPath, "web_portal/index.html"))
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
