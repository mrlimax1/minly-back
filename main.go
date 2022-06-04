package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"minly-back/src/checkers"
	"minly-back/src/database"
	"minly-back/src/utils"
	"net/http"
	"strconv"
)

var err = godotenv.Load()

var db = database.Connect()
var Site = new(database.Sites)

var getLink = func(w http.ResponseWriter, r *http.Request) {
	// Get INITIAL_LINK
	inLink := checkers.GetInitialLink(w, r)
	if checkers.CheckURLIsValid(inLink) == false {
		_, err := fmt.Fprint(w, "link is invalid")
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	if inLink[len(inLink)-1] != '/' {
		inLink += "/"
	}
	if checkers.CheckStatusLink(inLink, w) != true {
		return
	}

	err := db.Model(Site).Where("initial_link = ?", inLink).Select()
	if err != nil {
		log.Println(err)
		// обработка без сайта в бд
		site := &database.Sites{
			InitialLink: inLink,
			Link:        utils.GetRandomString(),
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
	_, err = db.Model(Site).
		Set("counter = counter + 1").
		Where("initial_link = ?", inLink).
		Update()
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprint(w, "{initial_link: "+Site.InitialLink+", link: "+Site.Link+", counter: "+strconv.Itoa(int(Site.Counter))+"}")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	http.HandleFunc("/get", getLink)
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
