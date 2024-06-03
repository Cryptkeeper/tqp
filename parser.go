package tqp

import (
	"regexp"
	"strings"
)

var regex = regexp.MustCompile(`(?i)([a-z]+):([a-z0-9\-_]+|"([\w()\-._ ]+)")+|([^ ]+)`)

// Find parses a string into a map of attributes, and a slice of noise words.
// Attributes must be in the format of key:value or key:"value". Keys must contain only letters a-z (case-insensitive),
// and are converted to lowercase when stored in the map.
// Unquoted values may contain letters, numbers, hyphens, and underscores.
// Quoted values may contain letters, numbers, hyphens, underscores, spaces, periods, and parentheses.
// Noise words are any words that are not formatted as attributes.
func Find(s string) (attrs map[string][]string, noise []string) {
	submatches := regex.FindAllStringSubmatch(s, -1)

	for _, submatch := range submatches {
		key := strings.ToLower(submatch[1])
		value := submatch[2]

		// if the value is quoted, use the unquoted submatch
		if n := submatch[3]; len(n) > 0 {
			value = n
		}

		// match unrelated words in the background to re-assemble query terms that aren't formatted attributes
		if n := submatch[4]; len(n) > 0 {
			noise = append(noise, n)
		}

		if len(key) > 0 && len(value) > 0 {
			if attrs == nil {
				attrs = make(map[string][]string)
			}
			attrs[key] = append(attrs[key], value)
		}
	}

	return
}
