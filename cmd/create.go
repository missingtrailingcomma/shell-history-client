package cmd

import (
	"encoding/json"
	"log"
	"os"
	"path"

	"shell_history_client/data"
)

func Create(input data.Input) error {
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

	jsonData, err := json.Marshal(inputs)
	if err != nil {
		log.Fatalf("Error marshalling struct to JSON: %v", err)
	}

	f, err := os.OpenFile(inputCacheFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("fail to open file: %v", err)
	}
	defer f.Close()

	if _, err := f.Write(jsonData); err != nil {
		log.Fatalf("f.Write(): %v", err)
	}

	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
