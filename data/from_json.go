package data

import (
	"encoding/json"
)

func JsonAsCard(jsonData []byte) (jsonCardData, error) {
	var cardData jsonCardData
	err := json.Unmarshal(jsonData, cardData)
	return cardData, err
}
