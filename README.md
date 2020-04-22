# Simple webhook - Jira -> Slack

# Configure at slack/slack.go
	jiraDomain  = "jira.com"
	jiraWebhook = "https://hooks.slack.com/services/XXXX/XXXXX/XXXX"

# Install
  git clone git@github.com:nopp/jira-webhook-slack.git
  
  go build
  
  ./jira-webhook-slack

# Jira webhook configuration
http://ipOfWebhookRunning:9229/alert

# Channel example of alert

<a href="https://imgbb.com/"><img src="https://i.ibb.co/ZSJdKcB/chan-Alert.png" alt="chan-Alert" border="0"></a>
