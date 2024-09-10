package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os/exec"

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

	// Run command.
	cmd := exec.Command(cmdData.Cmd)

	err = cmd.Start()

	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf("Executed command: '%s'.\n", cmdData.Cmd)

	err = cmd.Wait()

	if err != nil {
		fmt.Println(err)
	}
}
