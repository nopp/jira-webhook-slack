Simple webhook - Jira -> Slack
==============================

<img src="https://github.com/nopp/jira-webhook-slack/workflows/Go/badge.svg"/>

Configure at config.json
========================
    {
        "jiradomain":"xxjiradomainxx",
        "slackwebhook":"xxslackwebhookxx",
        "ipport":"xxiportxx"
    }
Install
=======
    git clone git@github.com:nopp/jira-webhook-slack.git
    go build
    ./jira-webhook-slack

Jira webhook configuration
==========================
    http://ipOfWebhookRunning:9229/alert


Running on docker
=================
    git clone https://github.com/nopp/jira-webhook-slack.git
    cd jira-webhook-slack/docker/
    docker build -t jira-webhook-slack:0.1 .

    docker run -d --name jira-webhook-slack \
    	-e "jiradomain=yourJiraDomain.com" \
    	-e "slackwebhook=https://hooks.slack.com/services/XXXX/XXXXX/XXXX" \
        -e "ipport=0.0.0.0:9229" \
    	-p 9229:9229 jira-webhook-slack:0.1

Channel example of alert
==========================
<a href="https://imgbb.com/"><img src="https://i.ibb.co/ZSJdKcB/chan-Alert.png" alt="chan-Alert" border="0"></a>
