package main

import (
	"fmt"
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/github"
)


func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	hook, _ := github.New(github.Options.Secret("daily-test-gihtub-secret-shane"))
	// hook, _ := github.New()

	fmt.Printf("%+v", r)
	payload, err := hook.Parse(r, github.PushEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			fmt.Printf("NOPE")
			// ok event wasn;t one of the ones asked to be parsed
		}
	}

	fmt.Printf("%+v", payload)
	switch payload.(type) {
		case github.ReleasePayload:
			release := payload.(github.ReleasePayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", release)

		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", pullRequest)
	}

}
