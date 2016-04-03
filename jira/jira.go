package jira

import (
	"fmt"

	"github.com/nlopes/slack"
)

type Jira struct {
	Config
}

type Config struct {
	Token   string `json:"token"`
	JiraUrl string `json:"jira_url"`
}

func (self Jira) Process(webHook WebHookEvent) {
	if self.isRequiredSendNotification(webHook) {
		params := NewParams(webHook)
		self.SendNotification(params)
	}
}

func (self Jira) isRequiredSendNotification(webHook WebHookEvent) bool {
	if webHook.Event == issueCreated {
		return true
	}

	if webHook.Event == issueUpdated {
		for _, field := range webHook.Changelog.Items {
			if field.Field == "status" {
				return true
			}
			if field.Field == "assignee" {
				return true
			}
		}
	}

	return false
}

func (self Jira) SendNotification(jiraParams Params) {
	slackParams := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Text:       "<" + self.Config.JiraUrl + jiraParams.Issue + "|*" + jiraParams.Summary + "*>",
		MarkdownIn: []string{"text", "pretext"},
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Assignee",
				Value: jiraParams.Assignee,
				Short: true,
			},
			slack.AttachmentField{
				Title: "Status",
				Value: jiraParams.Status,
				Short: true,
			},
			//			slack.AttachmentField{
			//				Title: "Creator",
			//				Value: jiraParams.Creator,
			//				Short: true,
			//			},
		},
	}
	slackParams.Attachments = []slack.Attachment{attachment}
	slackParams.IconURL = "https://a.slack-edge.com/ae7f/plugins/jira/assets/service_512.png" //"http://4.bp.blogspot.com/-cuf_fBZSARQ/T_68G2M6JFI/AAAAAAAAA1c/9zHj_fkeXds/s1600/Jackie+chan.jpg" //"http://risovach.ru/thumb/upload/200s400/2012/12/generator/ura_6457348_orig_.jpeg?bsx2q"
	slackParams.Username = "Jira"

	self.sendMessage(jiraParams.Modifier+" "+jiraParams.Action+" "+jiraParams.IssueType+" "+jiraParams.Issue, slackParams)
}

func (self Jira) sendMessage(message string, params slack.PostMessageParameters) {
	api := slack.New(self.Config.Token)

	channelID, timestamp, err := api.PostMessage("#test_go", message, params)
	if err != nil {
		fmt.Println("%s\n", err)
		return
	}
	fmt.Println("Message successfully sent to channel %s at %s", channelID, timestamp)
}
