package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("../../../"))
	http.Handle("/", fs)
	http.HandleFunc("/testsmartform", sayhello)
	http.HandleFunc("/getjson", getjson)
	http.ListenAndServe("0.0.0.0:8004", nil)
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte("I Hava Get Your Message:" + string(b)))
}

func getjson(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadFile("scheme.json")
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(body)
}
