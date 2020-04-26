package main

import (
	"jira-webhook-slack/conf"
	"jira-webhook-slack/slack"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	confInfo := conf.LoadConfiguration()
	router := mux.NewRouter()
	router.HandleFunc("/alert", slack.SendMessageToChannel).Methods("POST")

	log.Fatal(http.ListenAndServe(confInfo.IPPort, router))
}
