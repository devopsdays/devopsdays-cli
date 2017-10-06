package helpers

import "strings"

// TwitterClean strips the @ symbol from a twitter username if it exists, and returns the string
func TwitterClean(twitter string) (twitterClean string) {

	twitterClean = strings.Replace(strings.TrimSpace(strings.ToLower(twitter)), "@", "", 10)
	return

}
