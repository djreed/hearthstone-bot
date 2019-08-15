package data

import (
	"encoding/json"
)

func ToJson(card Card) ([]byte, error) {
	marshalled, marshalErr := json.Marshal(card)
	return marshalled, marshalErr
}
