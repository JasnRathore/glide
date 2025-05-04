package utils

import (
"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"	
	
	models "glide/models"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func RunCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin // Optional: allows interaction like prompts
	return cmd.Run()
}

func StartCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin // Optional: allows interaction like prompts
	return cmd.Start()
}

func StartCommandWithContext(ctx context.Context, command string, args ...string) (*exec.Cmd, error) {
	cmd := exec.CommandContext(ctx, command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd, cmd.Start()
}

func SetupSignalHandler(cancel context.CancelFunc, cmds ...*exec.Cmd) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		cancel()
		// Kill all commands
		for _, cmd := range cmds {
			if cmd != nil && cmd.Process != nil {
				cmd.Process.Kill() // Cross-platform way to kill process
			}
		}
	}()
}

func StructToJSON(data interface{}) ([]byte, error) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func WriteJSONToFile(filename string, jsonData []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}


func ReadProjectDetails() (models.ProjectDetails, error) {
	var project models.ProjectDetails

	filePath := "glide.config.json"
	
	// Read JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return project, err
	}
	defer file.Close()

	// Read file contents
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return project, err
	}

	// Unmarshal JSON into struct
	err = json.Unmarshal(byteValue, &project)
	if err != nil {
		return project, err
	}

	return project, nil
}