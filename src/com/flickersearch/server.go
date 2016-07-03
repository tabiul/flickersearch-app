package main

import (
	"com/flickersearch/authentication"
	"com/flickersearch/flicker"
	"com/flickersearch/history"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var apiKey string

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func handleAuthentication(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling query %s", r.URL.String())
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		u := user{}
		err := decoder.Decode(&u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintln(err)))
			return
		}

		err = authentication.CheckValidUser(u.Username, u.Password)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprint(err)))
			return
		}
		w.WriteHeader(http.StatusOK)

	}

}
func handleUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling query %s", r.URL.String())
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		u := user{}
		err := decoder.Decode(&u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintln(err)))
			return
		}
		err = authentication.AddUser(u.Username, u.Password)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprint(err)))
			return
		}
		w.WriteHeader(http.StatusCreated)
	}

}

func handleImage(w http.ResponseWriter, r *http.Request) {
	//only handle get for now
	if r.Method == "GET" {
		log.Printf("handling query %s", r.URL.String())
		parameters, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			log.Print(err)
			return
		}
		var search string
		var username string
		var page uint64 = 1

		username = parameters.Get("username")

		if search = parameters.Get("search"); search == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("search condition is missing"))
			return
		}

		if parameters.Get("page") != "" {
			page, err = strconv.ParseUint(parameters.Get("page"), 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("page must be a number"))
				return
			}
		}
		bytes, err := flicker.GetAllImages(apiKey, search, 5, page)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprint(err)))
			return
		}
		//update the history
		if username != "" {
			history.AddHistory(username, search)
		}
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Content-Length", strconv.Itoa(len(bytes)))
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}

}

func handleHistory(w http.ResponseWriter, r *http.Request) {
	//only handle get for now
	if r.Method == "GET" {
		log.Printf("handling query %s", r.URL.String())
		parameters, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			log.Print(err)
			return
		}
		var username string

		if username = parameters.Get("username"); username == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("search condition is missing"))
			return
		}

		bytes, err := history.GetHistory(username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprint(err)))
			return
		}
		if bytes != nil {
			w.Header().Add("Content-Type", "application/json")
			w.Header().Add("Content-Length", strconv.Itoa(len(bytes)))
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.WriteHeader(http.StatusOK)
			w.Write(bytes)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}

}

func main() {
	var port string
	var webAppDir string
	flag.StringVar(&port, "port", "8080", "specify the port to listen to")
	flag.StringVar(&apiKey, "apiKey", "", "specify the API key(required)")
	flag.StringVar(&webAppDir, "webapp", "", "specify the path to webapp(required)")
	flag.Parse()

	if apiKey == "" {
		log.Fatal("api key is required")
	}
	if webAppDir == "" {
		log.Fatal("webapp directory is required")
	}
	if _, err := os.Stat(webAppDir); err != nil {
		log.Fatalf("webapp directory %s is not valid", webAppDir)
	}
	http.HandleFunc("/authenticate", handleAuthentication)
	http.HandleFunc("/user", handleUser)
	http.HandleFunc("/image", handleImage)
	http.HandleFunc("/history", handleHistory)
	http.Handle("/", http.FileServer(http.Dir(webAppDir)))
	http.ListenAndServe(":"+port, nil)

}
