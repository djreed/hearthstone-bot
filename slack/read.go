package slack

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	bnet "github.com/djreed/hearthstone-bot/battlenet"
	ssl "github.com/djreed/hearthstone-bot/ssl"
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
		slack.OptionHTTPClient(ssl.HTTPSClient()),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)

	rtm := api.NewRTM(
		slack.RTMOptionDialer(ssl.Dialer()),
	)

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
				for _, matches := range captureMatcher.FindAllStringSubmatch(text, -1) {
					log.Printf("%s: Querying for '%s'", ev.User, matches[1])
					attachment := m.handleQuery(ev, matches[1])
					m.api.SendMessage(ev.Channel,
						slack.MsgOptionAttachments(
							attachment,
						),
					)
				}
			}

		case *slack.PresenceChangeEvent:
			log.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			fmt.Printf("Current latency: %v\n", ev.Value)

		case *slack.RTMError:
			log.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			log.Fatalf("Invalid credentials")
			return

		default:
			// Ignore other events..
			// log.Printf("Unexpected Event: %v -- %v\n", msg.Type, msg.Data)

		}
	}
}

func (m *slackManager) handleQuery(ev *slack.MessageEvent, query string) slack.Attachment {
	log.Printf("Searching Hearthstone data for '%s'", query)
	searchResult, _ /*res*/, _ /*err*/ := m.client.Hearthstone().Cards(query)

	if searchResult.CardCount < 1 {
		message := fmt.Sprintf("No results found for '%s'", query)
		log.Printf("%s: %s", ev.User, message)
		return slack.Attachment{
			Text: message,
		}
	} else if searchResult.CardCount > 1 {
		for _, card := range searchResult.Cards {
			if strings.EqualFold(card.Name, query) {
				log.Printf("%s: found '%s'", ev.User, card.Name)
				return cardAsAttachment(card)
			}
		}
		message := fmt.Sprintf("More than one result found for '%s'", query)
		log.Printf("%s: %s", ev.User, message)
		return slack.Attachment{Text: message}
	} else {
		card := searchResult.Cards[0]
		log.Printf("%s: found '%s'", ev.User, card.Name)
		return cardAsAttachment(card)
	}
}

func cardAsAttachment(card bnet.CardData) slack.Attachment {
	return slack.Attachment{
		Title:     formatCardTitle(card),
		TitleLink: card.CardURL(),
		Text:      formatCardString(card),
		ThumbURL:  card.Image,
	}
}

func formatCardTitle(card bnet.CardData) string {
	return fmt.Sprintf("{%d} %s", card.ManaCost, card.Name)
}

func formatCardString(card bnet.CardData) string {
	switch card.CardTypeID {
	case 4: // Minion
		return strings.Join([]string{
			fmt.Sprintf("%s", fixBoldText(card.Text)),
			fmt.Sprintf("_%s_", card.Flavor),
			fmt.Sprintf("*%d/%d*", card.Attack, card.Health),
		}, "\n")

	case 5: // Spell
		return strings.Join([]string{
			fmt.Sprintf("%s", fixBoldText(card.Text)),
			fmt.Sprintf("_%s_", card.Flavor),
		}, "\n")

	case 3: // Hero
		return strings.Join([]string{
			fmt.Sprintf("%s", fixBoldText(card.Text)),
			fmt.Sprintf("_%s_", card.Flavor),
		}, "\n")

	case 7: // Weapon
		return strings.Join([]string{
			fmt.Sprintf("%s", fixBoldText(card.Text)),
			fmt.Sprintf("_%s_", card.Flavor),
			fmt.Sprintf("*%d/%d*", card.Attack, card.Durability),
		}, "\n")

	default:
		log.Printf("UNKNOWN CARD TYPE FOR %s: %d", card.Name, card.CardTypeID)
	}
	return ""

}

func fixBoldText(text string) string {
	return strings.Replace(html2md.Convert(text), "**", "*", -1)
}
