package main

import (
	"log"
	"os"
)

func getBlizzardID() string {
	if BlizzardClientID := os.Getenv("BLIZZARD_ID"); BlizzardClientID != "" {
		log.Println("Found Blizzard ID")
		return BlizzardClientID
	}
	log.Fatal("Failed to find Blizzard ID in environment, exiting")
	os.Exit(1)
	return ""
}

func getBlizzardSecret() string {
	if blizzardClientSecret := os.Getenv("BLIZZARD_SECRET"); blizzardClientSecret != "" {
		log.Println("Found Blizzard Secret")
		return blizzardClientSecret
	}
	log.Print("Failed to find Blizzard Secret in environment, exiting")
	os.Exit(1)
	return ""
}

func getSlackToken() string {
	if slackToken := os.Getenv("SLACK_TOKEN"); slackToken != "" {
		log.Println("Found Slack Token")
		return slackToken
	}
	log.Print("Failed to find Slack Token in environment, exiting")
	os.Exit(1)
	return ""
}
