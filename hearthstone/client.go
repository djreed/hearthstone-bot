package hearthstone

import "github.com/djreed/hearthstone-bot/data"

type HearthClient interface {
	GetCard(string) data.Card
}

type HearthClientImpl struct {
	apiToken string `json:"-"`
}

func CreateClient(token string) *HearthClientImpl {
	client := &HearthClientImpl{
		apiToken: token,
	}

	return client
}

func (client *HearthClientImpl) GetCard(card string) data.Card {
	return nil
}

func (client *HearthClientImpl) GetCards(query string) []data.Card {
	return nil
}
