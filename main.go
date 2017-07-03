package main

import (
	"log"
	"net/http"
	"os"

	ih "./customutil"
)

var rootdir string

func _LogError(err error) {
	if err != nil {
		log.Fatal("Program Error : " + err.Error())
		os.Exit(1)
		panic(err)
	}
}

func main() {
	var iheader ih.IndexHandler
	rootdir, _ = os.Getwd()
	_, err := os.Stat(rootdir + "\\htdoc")
	_LogError(err)
	http.Handle("/", iheader)
	serverr := http.ListenAndServe(":8080", nil)
	_LogError(serverr)
}
