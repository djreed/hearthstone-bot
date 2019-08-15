package slack

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	bnet "github.com/djreed/hearthstone-bot/battlenet"
	"github.com/lunny/html2md"
	"github.com/nlopes/slack"
)

const (
	CAPTURE_REGEX = "\\[\\[([^\\[\\]]+)\\]\\]"
)

type slackManager struct {
	client *bnet.Client
	api    *slack.Client
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
		api:    api,
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
				log.Printf("%s: Querying for '%s'", ev.Username, query)
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

	if searchResult.CardCount < 1 {
		message := fmt.Sprintf("No results found for '%s'", query)
		log.Printf("%s: %s", ev.Username, message)
		m.api.SendMessage(ev.Channel,
			slack.MsgOptionText(message, false),
		)
	} else if searchResult.CardCount > 1 {
		message := fmt.Sprintf("More than one result found for '%s'", query)
		log.Printf("%s: %s", ev.Username, message)
		m.api.SendMessage(ev.Channel,
			slack.MsgOptionText(message, false),
		)
	} else {
		card := searchResult.Cards[0]
		log.Printf("%s: found '%s'", ev.Username, card.Name)
		m.api.SendMessage(ev.Channel,
			slack.MsgOptionAttachments(
				slack.Attachment{
					Text:     FormatCardString(card),
					ThumbURL: card.Image,
				},
			),
		)
	}
}

func FormatCardString(card bnet.CardData) string {
	switch card.CardTypeID {
	case 4: // Minion
		return strings.Join([]string{
			fmt.Sprintf("{%d} *%s*", card.ManaCost, card.Name),
			fmt.Sprintf("%s", html2md.Convert(card.Text)),
			fmt.Sprintf("_%s_", card.Flavor),
			fmt.Sprintf("*%d/%d*", card.Attack, card.Health),
		}, "\n")

	case 5: // Spell
		return strings.Join([]string{
			fmt.Sprintf("{%d} *%s*", card.ManaCost, card.Name),
			fmt.Sprintf("%s", html2md.Convert(card.Text)),
			fmt.Sprintf("_%s_", card.Flavor),
		}, "\n")

	case 3: // Hero
	case 7: // Weapon
	default:
		log.Printf("UNKNOWN CARD TYPE FOR %s: %d", card.Name, card.CardTypeID)
	}
	return ""

}
