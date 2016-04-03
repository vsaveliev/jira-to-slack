package main

import (
	"encoding/json"
	"github.com/vsaveliev/jira-to-slack/jira"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Port  string `json:"port"`
	Token string `json:"token"`

	JiraConfig jira.Config `json:"jira"`
}

func main() {
	config := GetConfig()

	jiraObject := jira.Jira{}
	jiraObject.Config = config.JiraConfig

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var webHook jira.WebHookEvent
		json.NewDecoder(r.Body).Decode(&webHook)

		go func() {
			jiraObject.Process(webHook)
		}()
	})

	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}

func GetConfig() Config {
	config := Config{}

	file, _ := os.Open("config.json")
	json.NewDecoder(file).Decode(&config)

	config.JiraConfig.Token = config.Token

	return config
}
