package slack

import (
	"fmt"
	"log"
	"os"
	"regexp"

	bnet "github.com/djreed/hearthstone-bot/battlenet"
	"github.com/nlopes/slack"
)

const (
	CAPTURE_REGEX = "\\[\\[([^\\[\\]]+)\\]\\]"
)

type slackManager struct {
	client *bnet.Client
	rtm    *slack.RTM
}

func NewManager(slackToken string, client *bnet.Client) *slackManager {
	api := slack.New(
		slackToken,
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)

	rtm := api.NewRTM()

	return &slackManager{
		client: client,
		rtm:    rtm,
	}
}

func (m *slackManager) ListenAndRespond() {
	go m.rtm.ManageConnection()

	for msg := range m.rtm.IncomingEvents {
		// fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello

		case *slack.ConnectedEvent:
			fmt.Println("Connected!")

		case *slack.MessageEvent:
			text := ev.Text
			captureMatcher := regexp.MustCompile(CAPTURE_REGEX)
			if captureMatcher.MatchString(text) {
				query := captureMatcher.FindStringSubmatch(text)[1]
				m.handleQuery(ev, query)
			}

		case *slack.PresenceChangeEvent:
			// log.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			// fmt.Printf("Current latency: %v\n", ev.Value)

		case *slack.RTMError:
			// log.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			// log.Fatalf("Invalid credentials")
			// return

		default:

			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
}

func (m *slackManager) handleQuery(ev *slack.MessageEvent, query string) {
	searchResult, _ /*res*/, _ /*err*/ := m.client.Hearthstone().Cards(query)

	var message string
	if searchResult.CardCount < 1 {
		message = "No results found"
	} else if searchResult.CardCount > 1 {
		message = fmt.Sprintf("Multiple results match '%s', be more specific", query)
	} else {
		message = searchResult.Cards[0].Name
	}
	m.rtm.SendMessage(m.rtm.NewOutgoingMessage(message, ev.Channel))
}
