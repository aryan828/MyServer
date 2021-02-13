package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aryan828/MyServer/logger"
)

var (
	certificate string = "/etc/letsencrypt/live/aayush.ninja/fullchain.pem"
	privateKey  string = "/etc/letsencrypt/live/aayush.ninja/privkey.pem"
	logs        *log.Logger
)

func redirect(w http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	logs.Println(req.RemoteAddr, "REDIRECTED")
	http.Redirect(w, req, target, http.StatusPermanentRedirect)
}

func handleBase(w http.ResponseWriter, r *http.Request) {
	if r.Host == "logs.aayush.ninja" {
		logFile, _ := ioutil.ReadFile(logger.LogFile)
		fmt.Fprint(w, string(logFile))
	} else {
		fmt.Fprintln(w, r.Host)
	}
	logs.Println(r.RemoteAddr)
}

func main() {
	logs = logger.InitializeLogging()

	go http.ListenAndServe(":80", http.HandlerFunc(redirect))

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleBase)
	log.Fatal(http.ListenAndServeTLS(":443", certificate, privateKey, mux))
}
