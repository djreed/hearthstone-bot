package main

import (
	slack "github.com/djreed/hearthstone-bot/slack"
)

var BlizzardClientID string
var BlizzardClientSecret string
var SlackToken string

func main() {
	slackToken := getSlackToken()
	slack.Start(slackToken)

	// query := getQuery()
	// blizzID := getBlizzardID()
	// blizzSecret := getBlizzardSecret()
	//
	// client := oauth.BlizzardOAuthClient(blizzID, blizzSecret)
	//
	// if data, _, err := client.Hearthstone().Cards(query); err != nil {
	// 	panic(err)
	// } else {
	// 	if dataStr, err := json.Marshal(data); err != nil {
	// 		panic(err)
	// 	} else {
	// 		log.Printf("Data: \n%s", dataStr)
	// 		os.Exit(0)
	// 	}
	// }
}
