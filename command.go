package main

import (
	"log"
	"os"
)

func getBlizzardID() string {
	if BlizzardClientID == "" {
		log.Print("Blizzard ID not provided in build env, searching runtime env")

		if BlizzardClientID = os.Getenv("BLIZZARD_ID"); BlizzardClientID == "" {
			log.Fatal("Failed to find Blizzard ID in environment, exiting")
			os.Exit(1)
		}
	}

	return BlizzardClientID
}

func getBlizzardSecret() string {
	if BlizzardClientSecret == "" {
		log.Print("Blizzard Secret not provided in build env, searching runtime env")

		if BlizzardClientSecret = os.Getenv("BLIZZARD_SECRET"); BlizzardClientSecret == "" {
			log.Print("Failed to find Blizzard Secret in environment, exiting")
			os.Exit(1)
		}
	}

	return BlizzardClientSecret
}

func getSlackToken() string {
	if SlackToken == "" {
		log.Print("Slack Token not provided in build env, searching runtime env")

		if SlackToken = os.Getenv("SLACK_TOKEN"); SlackToken == "" {
			log.Print("Failed to find Slack Token in environment, exiting")
			os.Exit(1)
		}
	}

	return SlackToken
}
