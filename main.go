package main

import (
	"github.com/vsaveliev/jira-to-slack/jira"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	jiraObject := jira.Jira{}
	jiraObject.Config = getConfig()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var webHook jira.WebHookEvent
		json.NewDecoder(r.Body).Decode(&webHook)

		go func() {
			jiraObject.Process(webHook)
		}()
	})

	log.Fatal(http.ListenAndServe(":7878", nil))
}

func getConfig() jira.Config{
	config := jira.Config{}

	file, _ := os.Open("config.json")
	json.NewDecoder(file).Decode(&config)

	return config
}
