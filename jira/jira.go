package jira

import (
	"fmt"

	"github.com/nlopes/slack"
)

const goodEventType = "good";
const warningEventType = "warning";
const dangerEventType = "danger";

type Jira struct {
	Config
}

type Config struct {
	Token     string `json:"token"`
	TicketUrl string `json:"ticket_url"`
	Channel   string `json:"channel"`
	BotName   string `json:"bot_name"`
	BotImgUrl string `json:"bot_img_url"`
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
			if field.Field == "status" || field.Field == "assignee" || field.Field == "priority" {
				return true
			}
		}
	}

	return false
}

func (self Jira) SendNotification(jiraParams Params) {
	slackParams := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Pretext: jiraParams.Modifier+" "+jiraParams.Action+" "+jiraParams.IssueType+" "+jiraParams.Issue,
		Text:       "<" + self.Config.TicketUrl + jiraParams.Issue + "|*" + jiraParams.Summary + "*>",
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
		Color: jiraParams.EventType,
	}
	slackParams.Attachments = []slack.Attachment{attachment}
	slackParams.IconURL = self.Config.BotImgUrl
	slackParams.Username = self.Config.BotName

	self.sendMessage("", slackParams)
}

func (self Jira) sendMessage(message string, params slack.PostMessageParameters) {
	api := slack.New(self.Config.Token)

	channelID, timestamp, err := api.PostMessage(self.Config.Channel, message, params)
	if err != nil {
		fmt.Println("%s\n", err)
		return
	}
	fmt.Println("Message successfully sent to channel %s at %s", channelID, timestamp)
}
