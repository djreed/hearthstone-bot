APPNAME="djreed/hearthstone-bot"
BINARY="hearthstone-bot"

default: run

docker:
	docker build -t ${APPNAME} -f Dockerfile .

build:
	go build -o ${BINARY} .

run: docker
	docker run --rm \
		-e BLIZZARD_ID \
		-e BLIZZARD_SECRET \
		-e SLACK_TOKEN \
		-it ${APPNAME}

heroku: docker
	heroku container:push bot

release: heroku
	heroku container:release bot
