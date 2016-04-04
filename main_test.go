package main_test

import (
	"github.com/vsaveliev/jira-to-slack"
	"github.com/vsaveliev/jira-to-slack/jira"
	"testing"
)

func TestCreatedIssueEvent(t *testing.T) {
	jiraObject := getJiraObject()
	event := getEventCreatedIssue()

	jiraObject.Process(event)
}

func TestUpdatedIssueEvent(t *testing.T) {
	jiraObject := getJiraObject()
	event := getEventUpdatedIssue()

	jiraObject.Process(event)
}

func TestResolvedIssueEvent(t *testing.T) {
	jiraObject := getJiraObject()
	event := getEventResolvedIssue()

	jiraObject.Process(event)
}

func getJiraObject() jira.Jira {
	config := main.GetConfig()

	jiraObject := jira.Jira{}
	jiraObject.Config = config.JiraConfig

	return jiraObject
}

func getEventCreatedIssue() jira.WebHookEvent {
	event := jira.WebHookEvent{}

	event.Event = "jira:issue_created"
	event.Issue.Fields.Assignee.DisplayName = "Vladislav Saveliev"
	event.Issue.Fields.Creator.DisplayName = "Michael Daddy"
	event.Issue.Fields.Reporter.DisplayName = "Michael Daddy"
	event.Modifier.DisplayName = "John Jones"
	event.Issue.Fields.IssueType.Name = "Bug"
	event.Issue.Fields.Status.Name = "Open"
	event.Issue.Fields.Summary = "Some summary of bug"
	event.Issue.Key = "OP-1111"

	return event
}

func getEventUpdatedIssue() jira.WebHookEvent {
	event := getEventCreatedIssue()

	event.Modifier.DisplayName = "Vladislav Saveliev"
	event.Event = "jira:issue_updated"
	items := make([]jira.ChangeLogItem, 2);
	items[0] = jira.ChangeLogItem{
		Field: "status",
		FromString: "Open",
		ToString: "In progress",
	}
	items[1] = jira.ChangeLogItem{
		Field: "assignee",
		FromString: "Vladislav Saveliev",
		ToString: "Daniel Smith",
	}
	event.Changelog.Items = items

	return event
}

func getEventResolvedIssue() jira.WebHookEvent {
	event := getEventCreatedIssue()

	event.Modifier.DisplayName = "Daniel Smith"
	event.Issue.Fields.Assignee.DisplayName = "Daniel Smith"
	event.Issue.Fields.Status.Name = "Resolved"
	event.Event = "jira:issue_updated"
	items := make([]jira.ChangeLogItem, 2);
	items[0] = jira.ChangeLogItem{
		Field: "status",
		FromString: "In progress",
		ToString: "Resolved",
	}
	event.Changelog.Items = items

	return event
}