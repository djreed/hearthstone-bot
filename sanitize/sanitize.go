package sanitize

import (
	"strings"
)

// Hearthstone's API is sensitive
func SanitizeString(query string) string {
	query = strings.Replace(query, "‘", "'", -1) // forward tick
	query = strings.Replace(query, "’", "'", -1) // back tick

	return query
}
