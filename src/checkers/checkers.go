package checkers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func checkURLIsValid(s string) (matched bool) {
	match, _ := regexp.MatchString(
		`^https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&\/=]*)$`,
		s)
	return match
}

func getInitialLink(w http.ResponseWriter, r *http.Request) string {
	initialLink, ok := r.URL.Query()["initial_link"]
	if !ok || len(initialLink) < 1 {
		_, err := fmt.Fprint(w, "initial_link not found")
		if err != nil {
			log.Fatal(err)
		}
	}
	return initialLink[0]
}
