# jira-to-slack
Integration jira with slack via web hooks.

Notifications about:
* Created issue
* Changed assignee
* Changed status

Example of work:
![ScreenShot](img:https://raw.github.com/vsaveliev/jira-to-slack/master/screenshot/2016-04-03_22-42-12.png)

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
