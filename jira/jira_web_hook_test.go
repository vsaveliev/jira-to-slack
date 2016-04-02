package jira_test

import (
	"github.com/vsaveliev/jira-to-slack/jira"
	"testing"

	"encoding/json"
	"strings"
)

var jsonCreatedIssueEvent string = `{
    "issue": {
        "fields": {
            "aggregateprogress": {
                "progress": 0,
                "total": 0
            },
            "aggregatetimeestimate": null,
            "aggregatetimeoriginalestimate": null,
            "aggregatetimespent": null,
            "assignee": {
                "active": true,
                "avatarUrls": {
                    "16x16": "https://jira.github.com/secure/useravatar?size=xsmall&ownerId=johnsmith&avatarId=10800",
                    "24x24": "https://jira.github.com/secure/useravatar?size=small&ownerId=johnsmith&avatarId=10800",
                    "32x32": "https://jira.github.com/secure/useravatar?size=medium&ownerId=johnsmith&avatarId=10800",
                    "48x48": "https://jira.github.com/secure/useravatar?ownerId=johnsmith&avatarId=10800"
                },
                "displayName": "Vladislav Saveliev",
                "emailAddress": "johnsmith@github.com",
                "key": "johnsmith",
                "name": "johnsmith",
                "self": "https://jira.github.com/rest/api/2/user?username=johnsmith",
                "timeZone": "Asia/Novosibirsk"
            },
            "attachment": [],
            "comment": {
                "comments": [],
                "maxResults": 0,
                "startAt": 0,
                "total": 0
            },
            "components": [
                {
                    "description": "Invoice lines, reports, balance, product renewal",
                    "id": "10209",
                    "name": "Billing",
                    "self": "https://jira.github.com/rest/api/2/component/10209"
                }
            ],
            "created": "2016-03-30T14:21:27.240+0200",
            "creator": {
                "active": true,
                "avatarUrls": {
                    "16x16": "https://jira.github.com/secure/useravatar?size=xsmall&avatarId=10122",
                    "24x24": "https://jira.github.com/secure/useravatar?size=small&avatarId=10122",
                    "32x32": "https://jira.github.com/secure/useravatar?size=medium&avatarId=10122",
                    "48x48": "https://jira.github.com/secure/useravatar?avatarId=10122"
                },
                "displayName": "Michael Daddy",
                "emailAddress": "maarten@github.com",
                "key": "maarten",
                "name": "maarten",
                "self": "https://jira.github.com/rest/api/2/user?username=maarten",
                "timeZone": "CET"
            },
            "customfield_10000": null,
            "customfield_10001": null,
            "customfield_10002": "9223372036854775807",
            "customfield_10006": null,
            "customfield_10007": null,
            "customfield_10100": null,
            "customfield_10200": null,
            "customfield_10201": null,
            "customfield_10202": null,
            "customfield_10300": "2|i010fl:",
            "customfield_10400": null,
            "customfield_10500": null,
            "customfield_10502": "0",
            "customfield_10504": null,
            "customfield_10505": null,
            "description": "Reseller reports to me thathem.",
            "duedate": null,
            "environment": null,
            "fixVersions": [],
            "issuelinks": [],
            "issuetype": {
                "avatarId": 10503,
                "description": "A problem which impairs or prevents the functions of the product.",
                "iconUrl": "https://jira.github.com/secure/viewavatar?size=xsmall&avatarId=10503&avatarType=issuetype",
                "id": "1",
                "name": "Bug",
                "self": "https://jira.github.com/rest/api/2/issuetype/1",
                "subtask": false
            },
            "labels": [],
            "lastViewed": null,
            "priority": {
                "iconUrl": "https://jira.github.com/images/icons/priorities/blocker.svg",
                "id": "1",
                "name": "Blocker",
                "self": "https://jira.github.com/rest/api/2/priority/1"
            },
            "progress": {
                "progress": 0,
                "total": 0
            },
            "project": {
                "avatarUrls": {
                    "16x16": "https://jira.github.com/secure/projectavatar?size=xsmall&pid=10000&avatarId=10011",
                    "24x24": "https://jira.github.com/secure/projectavatar?size=small&pid=10000&avatarId=10011",
                    "32x32": "https://jira.github.com/secure/projectavatar?size=medium&pid=10000&avatarId=10011",
                    "48x48": "https://jira.github.com/secure/projectavatar?pid=10000&avatarId=10011"
                },
                "id": "10000",
                "key": "OP",
                "name": "openprovider",
                "self": "https://jira.github.com/rest/api/2/project/10000"
            },
            "reporter": {
                "active": true,
                "avatarUrls": {
                    "16x16": "https://jira.github.com/secure/useravatar?size=xsmall&avatarId=10122",
                    "24x24": "https://jira.github.com/secure/useravatar?size=small&avatarId=10122",
                    "32x32": "https://jira.github.com/secure/useravatar?size=medium&avatarId=10122",
                    "48x48": "https://jira.github.com/secure/useravatar?avatarId=10122"
                },
                "displayName": "Michael Daddy",
                "emailAddress": "maarten@github.com",
                "key": "maarten",
                "name": "maarten",
                "self": "https://jira.github.com/rest/api/2/user?username=maarten",
                "timeZone": "CET"
            },
            "resolution": null,
            "resolutiondate": null,
            "status": {
                "description": "",
                "iconUrl": "https://jira.github.com/images/icons/statuses/generic.png",
                "id": "1",
                "name": "Open",
                "self": "https://jira.github.com/rest/api/2/status/1",
                "statusCategory": {
                    "colorName": "blue-gray",
                    "id": 2,
                    "key": "new",
                    "name": "To Do",
                    "self": "https://jira.github.com/rest/api/2/statuscategory/2"
                }
            },
            "subtasks": [],
            "summary": "Double invoicelines for renewals",
            "timeestimate": null,
            "timeoriginalestimate": null,
            "timespent": null,
            "timetracking": {},
            "updated": "2016-03-30T14:21:27.240+0200",
            "versions": [],
            "votes": {
                "hasVoted": false,
                "self": "https://jira.github.com/rest/api/2/issue/OP-8534/votes",
                "votes": 0
            },
            "watches": {
                "isWatching": false,
                "self": "https://jira.github.com/rest/api/2/issue/OP-8534/watchers",
                "watchCount": 0
            },
            "worklog": {
                "maxResults": 20,
                "startAt": 0,
                "total": 0,
                "worklogs": []
            },
            "workratio": -1
        },
        "id": "20951",
        "key": "OP-8534",
        "self": "https://jira.github.com/rest/api/2/issue/20951"
    },
    "timestamp": 1459340487369,
    "user": {
        "active": true,
        "avatarUrls": {
            "16x16": "https://jira.github.com/secure/useravatar?size=xsmall&avatarId=10122",
            "24x24": "https://jira.github.com/secure/useravatar?size=small&avatarId=10122",
            "32x32": "https://jira.github.com/secure/useravatar?size=medium&avatarId=10122",
            "48x48": "https://jira.github.com/secure/useravatar?avatarId=10122"
        },
        "displayName": "Michael Daddy",
        "emailAddress": "maarten@github.com",
        "key": "maarten",
        "name": "maarten",
        "self": "https://jira.github.com/rest/api/2/user?username=maarten",
        "timeZone": "CET"
    },
    "webhookEvent": "jira:issue_created"
}`

var jsonUpdatedIssueEvent string = `
{
    "changelog": {
        "id": "119101",
        "items": [
            {
                "field": "status",
                "fieldtype": "jira",
                "from": "1",
                "fromString": "Open",
                "to": "3",
                "toString": "In Progress"
            }
        ]
    },
    "issue": {
        "fields": {
            "aggregateprogress": {
                "progress": 0,
                "total": 0
            },
            "aggregatetimeestimate": null,
            "aggregatetimeoriginalestimate": null,
            "aggregatetimespent": null,
            "assignee": {
                "active": true,
                "avatarUrls": {
                    "16x16": "https://jira.github.com/secure/useravatar?size=xsmall&ownerId=vsaveliev&avatarId=10600",
                    "24x24": "https://jira.github.com/secure/useravatar?size=small&ownerId=vsaveliev&avatarId=10600",
                    "32x32": "https://jira.github.com/secure/useravatar?size=medium&ownerId=vsaveliev&avatarId=10600",
                    "48x48": "https://jira.github.com/secure/useravatar?ownerId=vsaveliev&avatarId=10600"
                },
                "displayName": "Vladislav Saveliev",
                "emailAddress": "vsaveliev@github.com",
                "key": "vsaveliev",
                "name": "vsaveliev",
                "self": "https://jira.github.com/rest/api/2/user?username=vsaveliev",
                "timeZone": "Asia/Novosibirsk"
            },
            "attachment": [
                {
                    "author": {
                        "active": true,
                        "avatarUrls": {
                            "16x16": "https://jira.github.com/secure/useravatar?size=xsmall&avatarId=10122",
                            "24x24": "https://jira.github.com/secure/useravatar?size=small&avatarId=10122",
                            "32x32": "https://jira.github.com/secure/useravatar?size=medium&avatarId=10122",
                            "48x48": "https://jira.github.com/secure/useravatar?avatarId=10122"
                        },
                        "displayName": "Michael Daddy",
                        "emailAddress": "maarten@github.com",
                        "key": "maarten",
                        "name": "maarten",
                        "self": "https://jira.github.com/rest/api/2/user?username=maarten",
                        "timeZone": "CET"
                    },
                    "content": "https://jira.github.com/secure/attachment/13361/Screen+Shot+2016-03-30+at+14.54.05.png",
                    "created": "2016-03-30T14:55:03.883+0200",
                    "filename": "Screen Shot 2016-03-30 at 14.54.05.png",
                    "id": "13361",
                    "mimeType": "image/png",
                    "self": "https://jira.github.com/rest/api/2/attachment/13361",
                    "size": 136477,
                    "thumbnail": "https://jira.github.com/secure/thumbnail/13361/_thumb_13361.png"
                }
            ],
            "comment": {
                "comments": [],
                "maxResults": 0,
                "startAt": 0,
                "total": 0
            },
            "components": [
                {
                    "id": "10201",
                    "name": "Domain management",
                    "self": "https://jira.github.com/rest/api/2/component/10201"
                }
            ],
            "created": "2016-03-30T14:21:27.240+0200",
            "creator": {
                "active": true,
                "avatarUrls": {
                    "16x16": "https://jira.github.com/secure/useravatar?size=xsmall&avatarId=10122",
                    "24x24": "https://jira.github.com/secure/useravatar?size=small&avatarId=10122",
                    "32x32": "https://jira.github.com/secure/useravatar?size=medium&avatarId=10122",
                    "48x48": "https://jira.github.com/secure/useravatar?avatarId=10122"
                },
                "displayName": "Michael Daddy",
                "emailAddress": "maarten@github.com",
                "key": "maarten",
                "name": "maarten",
                "self": "https://jira.github.com/rest/api/2/user?username=maarten",
                "timeZone": "CET"
            },
            "customfield_10000": null,
            "customfield_10001": null,
            "customfield_10002": "9223372036854775807",
            "customfield_10006": null,
            "customfield_10007": null,
            "customfield_10100": null,
            "customfield_10200": null,
            "customfield_10201": null,
            "customfield_10202": null,
            "customfield_10300": "2|i010fl:",
            "customfield_10400": null,
            "customfield_10500": null,
            "customfield_10502": "0",
            "customfield_10504": null,
            "customfield_10505": {
                "id": "10114",
                "self": "https://jira.github.com/rest/api/2/customFieldOption/10114",
                "value": "Expirations ( Vladislav Saveliev )"
            },
            "description": "Reseller reports to me that most of the domains that were renewed today were invoiced double.",
            "duedate": null,
            "environment": null,
            "fixVersions": [],
            "issuelinks": [],
            "issuetype": {
                "avatarId": 10503,
                "description": "A problem which impairs or prevents the functions of the product.",
                "iconUrl": "https://jira.github.com/secure/viewavatar?size=xsmall&avatarId=10503&avatarType=issuetype",
                "id": "1",
                "name": "Bug",
                "self": "https://jira.github.com/rest/api/2/issuetype/1",
                "subtask": false
            },
            "labels": [],
            "lastViewed": "2016-03-30T15:25:53.013+0200",
            "priority": {
                "iconUrl": "https://jira.github.com/images/icons/priorities/blocker.svg",
                "id": "1",
                "name": "Blocker",
                "self": "https://jira.github.com/rest/api/2/priority/1"
            },
            "progress": {
                "progress": 0,
                "total": 0
            },
            "project": {
                "avatarUrls": {
                    "16x16": "https://jira.github.com/secure/projectavatar?size=xsmall&pid=10000&avatarId=10011",
                    "24x24": "https://jira.github.com/secure/projectavatar?size=small&pid=10000&avatarId=10011",
                    "32x32": "https://jira.github.com/secure/projectavatar?size=medium&pid=10000&avatarId=10011",
                    "48x48": "https://jira.github.com/secure/projectavatar?pid=10000&avatarId=10011"
                },
                "id": "10000",
                "key": "OP",
                "name": "openprovider",
                "self": "https://jira.github.com/rest/api/2/project/10000"
            },
            "reporter": {
                "active": true,
                "avatarUrls": {
                    "16x16": "https://jira.github.com/secure/useravatar?size=xsmall&avatarId=10122",
                    "24x24": "https://jira.github.com/secure/useravatar?size=small&avatarId=10122",
                    "32x32": "https://jira.github.com/secure/useravatar?size=medium&avatarId=10122",
                    "48x48": "https://jira.github.com/secure/useravatar?avatarId=10122"
                },
                "displayName": "Michael Daddy",
                "emailAddress": "maarten@github.com",
                "key": "maarten",
                "name": "maarten",
                "self": "https://jira.github.com/rest/api/2/user?username=maarten",
                "timeZone": "CET"
            },
            "resolution": null,
            "resolutiondate": null,
            "status": {
                "description": "",
                "iconUrl": "https://jira.github.com/images/icons/statuses/generic.png",
                "id": "3",
                "name": "In Progress",
                "self": "https://jira.github.com/rest/api/2/status/3",
                "statusCategory": {
                    "colorName": "yellow",
                    "id": 4,
                    "key": "indeterminate",
                    "name": "In Progress",
                    "self": "https://jira.github.com/rest/api/2/statuscategory/4"
                }
            },
            "subtasks": [],
            "summary": "Double invoicelines for renewals",
            "timeestimate": null,
            "timeoriginalestimate": null,
            "timespent": null,
            "timetracking": {},
            "updated": "2016-03-30T15:25:53.083+0200",
            "versions": [
                {
                    "archived": false,
                    "description": "support issues from LIVE",
                    "id": "10200",
                    "name": "LIVE",
                    "releaseDate": "2020-05-29",
                    "released": false,
                    "self": "https://jira.github.com/rest/api/2/version/10200"
                }
            ],
            "votes": {
                "hasVoted": false,
                "self": "https://jira.github.com/rest/api/2/issue/OP-8534/votes",
                "votes": 0
            },
            "watches": {
                "isWatching": false,
                "self": "https://jira.github.com/rest/api/2/issue/OP-8534/watchers",
                "watchCount": 1
            },
            "worklog": {
                "maxResults": 20,
                "startAt": 0,
                "total": 0,
                "worklogs": []
            },
            "workratio": -1
        },
        "id": "20951",
        "key": "OP-8534",
        "self": "https://jira.github.com/rest/api/2/issue/20951"
    },
    "timestamp": 1459344353085,
    "user": {
        "active": true,
        "avatarUrls": {
            "16x16": "https://jira.github.com/secure/useravatar?size=xsmall&ownerId=vsaveliev&avatarId=10600",
            "24x24": "https://jira.github.com/secure/useravatar?size=small&ownerId=vsaveliev&avatarId=10600",
            "32x32": "https://jira.github.com/secure/useravatar?size=medium&ownerId=vsaveliev&avatarId=10600",
            "48x48": "https://jira.github.com/secure/useravatar?ownerId=vsaveliev&avatarId=10600"
        },
        "displayName": "Vladislav Saveliev",
        "emailAddress": "vsaveliev@github.com",
        "key": "vsaveliev",
        "name": "vsaveliev",
        "self": "https://jira.github.com/rest/api/2/user?username=vsaveliev",
        "timeZone": "Asia/Novosibirsk"
    },
    "webhookEvent": "jira:issue_updated"
}
`

func TestUnmarshalCreateIssue(t *testing.T) {
	event := jira.WebHookEvent{}

	dec := json.NewDecoder(strings.NewReader(jsonCreatedIssueEvent))

	err := dec.Decode(&event)
	if err != nil {
		t.Error("Create error:", err.Error())
	}

	if event.Event != "jira:issue_created" {
		t.Error("Incorrect parsing of event ", event.Event)
	}

	checkWebHookEventFields(t, event)
}

func TestUnmarshalUpdatedIssue(t *testing.T) {
	event := jira.WebHookEvent{}

	dec := json.NewDecoder(strings.NewReader(jsonUpdatedIssueEvent))

	err := dec.Decode(&event)
	if err != nil {
		t.Error("Create error:", err.Error())
	}

	if event.Event != "jira:issue_updated" {
		t.Error("Incorrect parsing of event ", event.Event)
	}

	item := event.Changelog.Items[0]

	if item.Field != "status" {
		t.Error("Incorrect parsing of changed field name ", item.Field)
	}

	if item.FromString != "Open" {
		t.Error("Incorrect parsing of old field value ", item.FromString)
	}

	if item.ToString != "In Progress" {
		t.Error("Incorrect parsing of new field value ", item.ToString)
	}

	checkWebHookEventFields(t, event)
}

func TestCreatedWebHookToParams(t *testing.T) {
	event := getWebHookEvent()

	jiraParams := jira.NewParams(*event)

	fields := event.Issue.Fields
	if jiraParams.Assignee != fields.Assignee.DisplayName {
		t.Error("Incorrect parsing of assignee name ", fields.Assignee.DisplayName, jiraParams.Assignee)
	}

	if jiraParams.Creator != fields.Creator.DisplayName {
		t.Error("Incorrect parsing of creator name ", fields.Creator.DisplayName, jiraParams.Creator)
	}

	if jiraParams.Reporter != fields.Reporter.DisplayName {
		t.Error("Incorrect parsing of reporter name ", fields.Reporter.DisplayName, jiraParams.Reporter)
	}

	if jiraParams.IssueType != fields.IssueType.Name {
		t.Error("Incorrect parsing of issue type ", fields.IssueType.Name, jiraParams.IssueType)
	}

	if jiraParams.Issue != event.Issue.Key {
		t.Error("Incorrect parsing of issue key ", event.Issue.Key, jiraParams.IssueType)
	}

	if jiraParams.Summary != fields.Summary {
		t.Error("Incorrect parsing of summary ", fields.Summary, jiraParams.Summary)
	}

	if jiraParams.Status != fields.Status.Name {
		t.Error("Incorrect parsing of status ", fields.Status.Name, jiraParams.Status)
	}

	if jiraParams.Action != "created" {
		t.Error("Incorrect parsing of action ", jiraParams.Action)
	}
}

func getWebHookEvent() *jira.WebHookEvent {
	event := &jira.WebHookEvent{}

	event.Event = "jira:issue_created"
	event.Issue.Fields.Assignee.DisplayName = "Vladislav Saveliev"
	event.Issue.Fields.Creator.DisplayName = "Michael Daddy"
	event.Issue.Fields.Reporter.DisplayName = "Michael Daddy"
	event.Issue.Fields.IssueType.Name = "Bug"
	event.Issue.Fields.Status.Name = "In Progress"
	event.Issue.Fields.Summary = "Double invoicelines for renewals"
	event.Issue.Key = "OP-1391"

	return event
}

func checkWebHookEventFields(t *testing.T, event jira.WebHookEvent) {
	fields := event.Issue.Fields
	if fields.Assignee.DisplayName != "Vladislav Saveliev" {
		t.Error("Incorrect parsing of assignee name ", fields.Assignee.DisplayName)
	}

	if fields.Creator.DisplayName != "Michael Daddy" {
		t.Error("Incorrect parsing of creator name ", fields.Creator.DisplayName)
	}

	if fields.IssueType.Name != "Bug" {
		t.Error("Incorrect parsing of issue type ", fields.IssueType.Name)
	}

	if event.Issue.Key != "OP-8534" {
		t.Error("Incorrect parsing of key ", event.Issue.Key)
	}

	if fields.Summary != "Double invoicelines for renewals" {
		t.Error("Incorrect parsing of summary ", fields.Summary)
	}

	if fields.Priority.Name != "Blocker" {
		t.Error("Incorrect parsing of priority ", fields.Priority.Name)
	}
}
