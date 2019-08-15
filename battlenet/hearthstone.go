package battlenet

import "net/url"

type HearthService struct {
	client *Client
}

type SearchResult struct {
	CardCount int        `json:"cardCount"`
	Cards     []CardData `json:"cards"`
	Page      int        `json:"page"`
	PageCount int        `json:"pageCount"`
}

type CardData struct {
	ID         int    `json:"id"`
	Attack     int    `json:"attack"`
	Health     int    `json:"health"`
	ManaCost   int    `json:"manaCost"`
	RarityID   int    `json:"rarityId"`
	Name       string `json:"name"`
	Text       string `json:"text"`
	Flavor     string `json:"flavorText"`
	Image      string `json:"image"`
	CardTypeID int    `json:"cardTypeId"`
}

func baseQuery() url.Values {
	base := url.Values{}
	base.Set("locale", "en_US")
	return base
}

func (s *HearthService) Cards(filter string) (*SearchResult, *Response, error) {
	req, err := s.client.NewRequest("GET", "hearthstone/cards", nil)

	query := baseQuery()
	query.Add("textFilter", filter)
	req.URL.RawQuery = query.Encode()

	if err != nil {
		return nil, nil, err
	}

	var cards SearchResult
	resp, err := s.client.Do(req, &cards)
	if err != nil {
		return nil, resp, err
	}

	return &cards, resp, nil
}
