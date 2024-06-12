package cmd

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
	"shell_history_client/data"
	"text/template"

	"github.com/bazelbuild/rules_go/go/tools/bazel"
)

type PageData struct {
	Title  string
	Inputs []data.Input
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

		inputs = append(inputs, input)
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
		data := PageData{
			Title:  "Gummy Bear 🍬🐻",
			Inputs: inputs,
		}

		// Execute the template with the data
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Start the HTTP server
	log.Println("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	return nil
}
