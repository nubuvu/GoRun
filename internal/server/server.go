package server

import (
	"GoRun/configs"
	"log"
	"net/http"
)

func rend(w http.ResponseWriter, msg string) {
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Println(err)
	}
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "imge")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "favicon")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "ping")
}

func robotsHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "robots")
}

func Run(conf configs.ConfI) {
	http.HandleFunc("/", imgHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/robots.txt", robotsHandler)

	log.Println("Server starting...")
	if err := http.ListenAndServe(":"+conf.GetPort(), nil); err != nil {
		log.Fatalln(err)
	}
}
