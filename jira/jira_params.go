package jira

const issueCreated string = "jira:issue_created"
const issueUpdated string = "jira:issue_updated"

type Params struct {
	Summary   string
	Reporter  string
	Assignee  string
	Creator   string
	Modifier  string
	Action    string
	Status    string
	Issue     string
	IssueType string
	Title     string
	Event     string
}

func NewParams(event WebHookEvent) Params {
	params := Params{
		Summary:   event.Issue.Fields.Summary,
		Assignee:  event.Issue.Fields.Assignee.DisplayName,
		Creator:   event.Issue.Fields.Creator.DisplayName,
		Reporter:  event.Issue.Fields.Reporter.DisplayName,
		IssueType: event.Issue.Fields.IssueType.Name,
		Issue:     event.Issue.Key,
		Status:    event.Issue.Fields.Status.Name,
		Modifier:  event.Modifier.DisplayName,
		Event:     event.Event,
	}

	switch {
	case params.Event == issueUpdated:
		params.Action = "updated"
		for _, field := range event.Changelog.Items {
			switch {
			case field.Field == "status":
				if field.ToString == "Resolved" {
					params.Action = "resolved"
					continue
				}
				params.Status = field.FromString + " --> " + field.ToString
			case field.Field == "assignee":
				params.Assignee = field.FromString + " --> " + field.ToString
			}
		}
	case params.Event == issueCreated:
		params.Action = "created"
	}

	return params
}
