BINARY="hearthstone"
LD_FLAGS=-ldflags "-X main.BlizzardClientID=$(BLIZZARD_ID) -X main.BlizzardClientSecret=$(BLIZZARD_SECRET)"

build:
	go build -o $(BINARY) $(LD_FLAGS) .
		