package slack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/slack-go/slack"
)

type issueJira struct {
	ID        int    `json:"id"`
	Timestamp string `json:"timestamp"`
	Issue     struct {
		Expand string `json:"expand"`
		ID     string `json:"id"`
		Self   string `json:"self"`
		Key    string `json:"key"`
		Fields struct {
			Summary     string   `json:"summary"`
			Created     string   `json:"created"`
			Description string   `json:"description"`
			Labels      []string `json:"labels"`
			Priority    string   `json:"priority"`
		} `json:"fields"`
	} `json:"issue"`
	User struct {
		Self         string `json:"self"`
		AccoundID    string `json:"accoundId"`
		EmailAddress string `json:"emailAddress"`
		AvatarUrls   struct {
			One6X16  string `json:"16x16"`
			Four8X48 string `json:"48x48"`
		} `json:"avatarUrls"`
		DisplayName string `json:"displayName"`
		Active      string `json:"active"`
	} `json:"user"`
	Changelog struct {
		Items []struct {
			To         interface{} `json:"to"`
			FromString string      `json:"fromString"`
			From       interface{} `json:"from"`
			Fieldtype  string      `json:"fieldtype"`
			Field      string      `json:"field"`
		} `json:"items"`
		ID int `json:"id"`
	} `json:"changelog"`
	Comment struct {
		Self   string `json:"self"`
		ID     string `json:"id"`
		Author struct {
			Self         string `json:"self"`
			AccountID    string `json:"accountId"`
			EmailAddress string `json:"emailAddress"`
			AvatarUrls   struct {
				One6X16  string `json:"16x16"`
				Four8X48 string `json:"48x48"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
		} `json:"author"`
		Body         string `json:"body"`
		UpdateAuthor struct {
			Self         string `json:"self"`
			AccoundID    string `json:"accoundId"`
			EmailAddress string `json:"emailAddress"`
			AvatarUrls   struct {
				One6X16  string `json:"16x16"`
				Four8X48 string `json:"48x48"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
		} `json:"updateAuthor"`
		Created string `json:"created"`
		Updated string `json:"updated"`
	} `json:"comment"`
	WebhookEvent string `json:"webhookEvent"`
}

// SendMessageToChannel Responsible for send msgs to channel
func SendMessageToChannel(w http.ResponseWriter, r *http.Request) {

	var issue issueJira

	_ = json.NewDecoder(r.Body).Decode(&issue)

	attachment := slack.Attachment{
		Color:      "good",
		Fallback:   "You successfully posted by Incoming Webhook URL!",
		AuthorName: issue.Issue.Key + " - " + issue.Issue.Fields.Summary,
		AuthorLink: "https://jira.preventsenior.io/browse/" + issue.Issue.Key,
		Text:       issue.User.DisplayName,
		Ts:         json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	// Webhook for channel #infra-test
	err := slack.PostWebhook("xxxCHANGEitxxx", &msg)
	if err != nil {
		fmt.Println(err)
	}
}
