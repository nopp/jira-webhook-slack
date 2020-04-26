# Simple webhook - Jira -> Slack

# Configure at conf/conf.go

    JiraDomain  = "jira.com"
    SlackWebhook = "https://hooks.slack.com/services/XXXX/XXXXX/XXXX"
    IPPort = "0.0.0.0:9229"

# Install
    git clone git@github.com:nopp/jira-webhook-slack.git
    go build
    ./jira-webhook-slack

# Jira webhook configuration
    http://ipOfWebhookRunning:9229/alert

# Channel example of alert

<a href="https://imgbb.com/"><img src="https://i.ibb.co/ZSJdKcB/chan-Alert.png" alt="chan-Alert" border="0"></a>
