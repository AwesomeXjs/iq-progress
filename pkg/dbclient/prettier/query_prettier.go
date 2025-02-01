package prettier

import (
	"fmt"
	"strconv"
	"strings"
)

// PlaceholderDollar and PlaceholderQuestion are SQL placeholder symbols.
const (
	PlaceholderDollar   = "$"
	PlaceholderQuestion = "?"
)

// Pretty formats a SQL query by replacing placeholders with corresponding argument values.
func Pretty(query string, placeholder string, args ...any) string {
	// Iterate through all arguments and replace placeholders in the query with their string representation.
	for i, param := range args {
		var value string
		// Convert the argument to a string representation.
		switch v := param.(type) {
		case string:
			value = fmt.Sprintf("%q", v)
		case []byte:
			value = fmt.Sprintf("%q", string(v))
		default:
			value = fmt.Sprintf("%v", v)
		}

		// Replace the placeholder with the value in the query.
		query = strings.Replace(query, fmt.Sprintf("%s%s", placeholder, strconv.Itoa(i+1)), value, -1)
	}

	// Clean up the query by removing tabs, newlines, and trimming any extra spaces.
	query = strings.ReplaceAll(query, "\t", "")
	query = strings.ReplaceAll(query, "\n", " ")

	return strings.TrimSpace(query)
}
