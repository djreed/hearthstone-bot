package data

// A single Hearthstone card
type Card interface {
	// The card name
	Name() string

	// The card's primary card text
	Text() string

	// The card's flavor text
	Flavor() string

	// The URL of the full-card image
	Image() string
}

type jsonCardData struct {
	CardName   string `json:"name"`
	CardText   string `json:"text"`
	FlavorText string `json:"flavor"`
	ImageURL   string `json:"image"`
}

func (card *jsonCardData) Name() string {
	return card.CardName
}

func (card *jsonCardData) Text() string {
	return card.CardText
}

func (card *jsonCardData) Flavor() string {
	return card.FlavorText
}

func (card *jsonCardData) Image() string {
	return card.ImageURL
}

func NewCard(name, text, flavor, image string) *jsonCardData {
	return &jsonCardData{
		CardName:   name,
		CardText:   text,
		FlavorText: flavor,
		ImageURL:   image,
	}
}
