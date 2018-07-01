package main

import (
	"./goBot"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Result struct {
	Website string `json:"website"`
	Option  string `json:"option"`
}

func linksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		res := Result{}
		err := json.NewDecoder(r.Body).Decode(&res)
		if err != nil {
			fmt.Println(err)
		}
		links, err := geturls.GetLinks(res.Website, "", "", 15, 1)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(links)
	}
}

func main() {
	http.HandleFunc("/LIVE", linksHandler)
	fmt.Println("Serving on localhost:8008/LIVE")
	log.Fatal(http.ListenAndServe(":8008", nil))
}
