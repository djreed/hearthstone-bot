package main

import (
	"log"
	"os"
)

func getQuery() string {
	if len(os.Args) <= 1 {
		log.Fatal("Invalid number of arguments: must provide card name")
	}
	queryArgument := os.Args[1]
	log.Printf("Query: %s", queryArgument)
	return queryArgument
}

func getBlizzardID() string {
	if BlizzardClientID == "" {
		log.Print("Blizzard ID not provided in build environment, searching inputs")

		if len(os.Args) > 2 {
			BlizzardClientID = os.Args[2]
		}

		if BlizzardClientID == "" {
			log.Print("No Blizzard ID found in arguments, searching environment")

			if BlizzardClientID = os.Getenv("BLIZZARD_ID"); BlizzardClientID == "" {
				log.Fatal("Failed to find Blizzard ID in environment, exiting")
				os.Exit(1)
			}
		}
	}

	return BlizzardClientID
}

func getBlizzardSecret() string {
	if BlizzardClientSecret == "" {
		log.Print("Blizzard Secret not provided in build environment, searching inputs")

		if len(os.Args) > 2 {
			BlizzardClientSecret = os.Args[2]
		}

		if BlizzardClientSecret == "" {
			log.Print("No Blizzard Secret found in arguments, searching environment")

			if BlizzardClientSecret = os.Getenv("BLIZZARD_SECRET"); BlizzardClientSecret == "" {
				log.Print("Failed to find Blizzard Secret in environment, exiting")
				os.Exit(1)
			}
		}
	}

	return BlizzardClientSecret
}

func getSlackToken() string {
	if SlackToken == "" {
		log.Print("Slack Token not provided in build environment, searching inputs")

		if len(os.Args) > 2 {
			SlackToken = os.Args[2]
		}

		if SlackToken == "" {
			log.Print("No Slack Token found in arguments, searching environment")

			if SlackToken = os.Getenv("SLACK_TOKEN"); SlackToken == "" {
				log.Print("Failed to find Slack Token in environment, exiting")
				os.Exit(1)
			}
		}
	}

	return BlizzardClientSecret
}
