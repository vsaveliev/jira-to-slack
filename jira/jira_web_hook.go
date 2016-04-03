package jira

type WebHookEvent struct {
	Event string `json:"webhookEvent"`

	Issue struct {
		Fields struct {
			Assignee  User `json:"assignee"`
			Creator   User `json:"creator"`
			IssueType struct {
				Name string `json:"name"`
			} `json:"issuetype"`
			Priority struct {
				Name string `json:"name"`
			} `json:"priority"`
			Reporter User   `json:"reporter"`
			Summary  string `json:"summary"`
			Status   struct {
				Name string `json:"name"`
			} `json:"status"`
		} `json:"fields"`
		Key string `json:"key"`
	} `json:"issue"`

	Modifier User `json:"user"`

	// for Updated issue
	Changelog struct {
		Items []ChangeLogItem `json:"items"`
	} `json:"changelog"`
}

type User struct {
	DisplayName string `json:"displayName"`
}

type ChangeLogItem struct {
	Field      string `json:"field"`
	FromString string `json:"fromString"`
	ToString   string `json:"toString"`
}
