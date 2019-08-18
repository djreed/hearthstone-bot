APPNAME="djreed/hearthstone-bot"
BINARY="hearthstone-bot"

default: run

docker: vendor
	docker build -t ${APPNAME} -f Dockerfile .

vendor:
	GO111MODULE=on go mod vendor

build: vendor
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
