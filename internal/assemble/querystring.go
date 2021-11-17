package assemble

import (
	"net/url"
	"strings"

	"github.com/jonaslu/ain/internal/pkg/data"
)

const defaultQueryDelim = "&"

func isHex(currentChar byte) bool {
	switch {
	case '0' <= currentChar && currentChar <= '9':
		return true
	case 'a' <= currentChar && currentChar <= 'f':
		return true
	case 'A' <= currentChar && currentChar <= 'F':
		return true
	}

	return false
}

// Borrowed from net/url in the go standard library
const upperhex = "0123456789ABCDEF"

func queryEscape(queryString string) string {
	var result strings.Builder
	result.Grow(len(queryString))

	for i := 0; i < len(queryString); i++ {
		currentChar := queryString[i]

		if 'a' <= currentChar && currentChar <= 'z' ||
			'A' <= currentChar && currentChar <= 'Z' ||
			'0' <= currentChar && currentChar <= '9' ||
			currentChar == '+' ||
			i+2 < len(queryString) && isHex(queryString[i+1]) && isHex(queryString[i+2]) {

			result.WriteByte(currentChar)
		} else {
			if currentChar == ' ' {
				result.WriteByte('+')
			} else {
				result.WriteByte('%')
				result.WriteByte(upperhex[currentChar>>4])
				result.WriteByte(upperhex[currentChar&15])
			}
		}
	}

	return result.String()
}

func addQueryString(host *url.URL, parse *data.Parse) {
	if host.RawQuery == "" && len(parse.Query) == 0 {
		return
	}

	queryDelim := defaultQueryDelim
	if parse.Config.QueryDelim != nil {
		queryDelim = *parse.Config.QueryDelim
	}

	queryString := host.RawQuery
	if host.RawQuery != "" && len(parse.Query) > 0 {
		queryString = queryString + queryDelim
	}

	queryString = queryString + strings.Join(parse.Query, queryDelim)

	if queryDelim == "" || !strings.Contains(queryString, queryDelim) {
		host.RawQuery = queryEscape(queryString)
		return
	}

	var encodedKeyValuePairs []string
	for _, keyValuePairStr := range strings.Split(queryString, queryDelim) {
		var encodedKeyValuePair string

		keyValuePair := strings.SplitN(keyValuePairStr, "=", 2)
		if len(keyValuePair) == 2 {
			encodedKeyValuePair = strings.Join(
				[]string{
					queryEscape(keyValuePair[0]),
					queryEscape(keyValuePair[1]),
				},
				"=",
			)
		} else {
			encodedKeyValuePair = queryEscape(keyValuePairStr)
		}

		encodedKeyValuePairs = append(encodedKeyValuePairs, encodedKeyValuePair)
	}

	host.RawQuery = strings.Join(encodedKeyValuePairs, queryDelim)
}