BINARY="hearthstone-bot"
LD_FLAGS=-ldflags "-X main.BlizzardClientID=$(BLIZZARD_ID) -X main.BlizzardClientSecret=$(BLIZZARD_SECRET) -X main.SlackToken=$(SLACK_TOKEN)"

build:
	go build -o $(BINARY) $(LD_FLAGS) .
		