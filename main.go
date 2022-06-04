package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// DOTENV
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	// DATABASE
	db := databaseConnect()
	// SERVER
	getLink := func(w http.ResponseWriter, r *http.Request) {
		// Get INITIAL_LINK
		inLink := getInitialLink(w, r)
		if checkURLIsValid(inLink) == false {
			_, err := fmt.Fprint(w, "link is invalid")
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		if inLink[len(inLink)-1] != '/' {
			inLink += "/"
		}

		site := new(Sites)
		err = db.Model(site).Where("initial_link = ?", inLink).Select()
		if err != nil {
			log.Println(err)
			// обработка без сайта в бд
			site = &Sites{
				InitialLink: inLink,
				Link:        getRandomString(),
				Counter:     1,
			}
			_, err = db.Model(site).Insert()
			if err != nil {
				log.Println(err)
			}
			_, err := fmt.Fprint(w, "{initial_link: "+site.InitialLink+", link: "+site.Link+", counter: "+strconv.Itoa(int(site.Counter))+"}")
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		_, err := fmt.Fprint(w, "{initial_link: "+site.InitialLink+", link: "+site.Link+", counter: "+strconv.Itoa(int(site.Counter))+"}")
		if err != nil {
			log.Fatal(err)
		}
	}
	http.HandleFunc("/get", getLink)
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
