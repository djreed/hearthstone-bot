package main

import (
	"encoding/json"
	"log"
	"os"

	oauth "github.com/djreed/hearthstone-bot/oauth"
)

var BlizzardClientID string
var BlizzardClientSecret string

func main() {
	// query := getQuery()
	id := getBlizzardID()
	secret := getBlizzardSecret()

	client := oauth.BlizzardOAuthClient(id, secret)

	if data, _, err := client.Hearthstone().Rafaam(); err != nil {
		panic(err)
	} else {
		if dataStr, err := json.Marshal(data); err != nil {
			panic(err)
		} else {
			log.Printf("Data: \n%s", dataStr)
			os.Exit(0)
		}
	}
}

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

	log.Println("Blizzard ID successfully found")

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

	log.Println("Blizzard Secret successfully found")

	return BlizzardClientSecret
}
