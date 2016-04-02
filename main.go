package main

import (
	"bitbucket.org/vsaveliev/slack/jira"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	jiraObject := jira.Jira{}
	jiraObject.Config = jira.Config{
		Token:   "xoxp-27287745891-27331526278-30312717234-ae3f8c0320",
		JiraUrl: "https://jira.openprovider.nl/browse/",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var webHook jira.WebHookEvent
		json.NewDecoder(r.Body).Decode(&webHook)

		go func() {
			jiraObject.Process(webHook)
		}()
	})

	log.Fatal(http.ListenAndServe(":7878", nil))
}
