package main

import (
	"jira-webhook-slack/slack"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/alert", slack.SendMessageToChannel).Methods("POST")

	log.Fatal(http.ListenAndServe("0.0.0.0:9229", router))
}
