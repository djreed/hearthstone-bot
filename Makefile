APPNAME="djreed/hearthstone-bot"
BINARY="hearthstone-bot"
LD_FLAGS=-ldflags "-X main.BlizzardClientID=$(BLIZZARD_ID) -X main.BlizzardClientSecret=$(BLIZZARD_SECRET) -X main.SlackToken=$(SLACK_TOKEN)"

default: run

docker:
	docker build -t ${APPNAME} -f Dockerfile .

build:
	go build -o $(BINARY) $(LD_FLAGS) .
		

run:
	docker run --rm \
		-e BLIZZARD_ID \
		-e BLIZZARD_SECRET \
		-e SLACK_TOKEN \
		-it ${APPNAME}
