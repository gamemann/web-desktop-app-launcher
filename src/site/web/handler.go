package web

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gamemann/web-desktop-app-launcher/config"
)

type CommandData struct {
	Cmd string `json:"command"`
}

func RootHandler(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		http.Error(w, "Error parsing template.", http.StatusInternalServerError)

		fmt.Println(err)

		return
	}

	err = tmpl.Execute(w, cfg.Apps)

	if err != nil {
		http.Error(w, "Error executing template.", http.StatusInternalServerError)

		fmt.Println(err)
	}
}

func BackendHandler(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	// Get type and application.
	if r.Method != "POST" {
		http.Error(w, "Wrong method.", http.StatusMethodNotAllowed)

		return
	}

	var cmdData CommandData

	err := json.NewDecoder(r.Body).Decode(&cmdData)

	if err != nil {
		http.Error(w, "Error decoding JSON data.", http.StatusInternalServerError)

		return
	}

	// We'll want to make sure we handle spaces properly.
	cmdSplit := strings.Fields(cmdData.Cmd)

	// Run command.
	cmd := exec.Command(cmdSplit[0], cmdSplit[1:]...)

	// Get current environment.
	env := os.Environ()

	// Add environmental variables (yeah we need a better way on handling this in the future).
	for _, app := range cfg.Apps {
		if app.Start == cmdData.Cmd || app.Stop == cmdData.Cmd {
			for k, v := range app.Env {
				env = append(env, fmt.Sprintf("%s=%s", k, v))
			}
		}
	}

	cmd.Env = env

	// We need to get pipes now for logging.
	outPipe, _ := cmd.StdoutPipe()
	errPipe, _ := cmd.StderrPipe()

	err = cmd.Start()

	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf("Executed command: '%s'.\n", cmdData.Cmd)

	if cmd.Process == nil {
		fmt.Println("Process doesn't exist.")

		return
	}

	// Handle logging.
	if cfg.Web.LogToFile {
		go func() {
			fName := fmt.Sprintf("%s/apps/%d.log", cfg.Web.LogDirectory, cmd.Process.Pid)

			logFile, err := os.Create(fName)

			if err != nil {
				fmt.Printf("Failed to create log for process with ID '%d' (%s)", cmd.Process.Pid, fName)
				fmt.Println(err)

				return
			}

			outWriter := bufio.NewWriter(logFile)
			errWriter := bufio.NewWriter(logFile)

			// Handle stdout writes.
			go func() {
				scanner := bufio.NewScanner(outPipe)

				for scanner.Scan() {
					line := scanner.Text()
					outWriter.WriteString(line + "\n")
					outWriter.Flush()
				}
			}()

			// Handle stderr writes.
			go func() {
				scanner := bufio.NewScanner(errPipe)

				for scanner.Scan() {
					line := scanner.Text()
					errWriter.WriteString(line + "\n")
					errWriter.Flush()
				}
			}()
		}()
	}
}
