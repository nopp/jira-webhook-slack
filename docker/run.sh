#!/bin/bash

if [ "x" == "x$jiradomain" ]; then
  echo "FAIL: jiradomain is not set"
  exit 1
fi

if [ "x" == "x$slackwebhook" ]; then
  echo "FAIL: slackwebhook is not set"
  exit 1
fi

if [ "x" == "x$ipport" ]; then
  echo "FAIL: ipport is not set"
  exit 1
fi

sed -i s/xxjiradomainxx/"$jiradomain"/ conf/conf.go
sed -i s/xxslackwebhookxx/"$slackwebhook"/ conf/conf.go
sed -i s/xxiportxx/"$ipport"/ conf/conf.go

./jira-webhook-slack