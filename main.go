package main

import (
	"github.com/djreed/hearthstone-bot/oauth"
	slack "github.com/djreed/hearthstone-bot/slack"
)

func main() {
	blizzID := getBlizzardID()
	blizzSecret := getBlizzardSecret()
	slackToken := getSlackToken()

	client := oauth.BlizzardOAuthClient(blizzID, blizzSecret)

	manager := slack.NewManager(slackToken, client)
	manager.ListenAndRespond()
}
