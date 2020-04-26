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

sed -i s/xxjiradomainxx/"$jiradomain"/ config.json
sed -i s/xxslackwebhookxx/"$slackwebhook"/ config.json
sed -i s/xxiportxx/"$ipport"/ config.json

./jira-webhook-slack