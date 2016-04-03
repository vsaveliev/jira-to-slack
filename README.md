# jira-to-slack
Integration jira with slack via web hooks.

Notifications about:
* Created issue
* Changed assignee
* Changed status

TODO:
* send direct notification to assignee
* more flexible config

**./config.json**:
```json
{
  "port": "7878",
  "token": "some token",

  "jira": {
    "ticket_url": "https://jira.github.com/browse/",
    "channel": "#somechannel",
    "bot_name": "Jira",
    "bot_img_url": "https://a.slack-edge.com/ae7f/plugins/jira/assets/service_512.png"
  }
}
```
