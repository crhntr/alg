package main

import (
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	cookie := http.Cookie{Name: "sid", Value: bson.NewObjectId().Hex(), Expires: time.Now().Add(365 * 24 * time.Hour)}

	http.SetCookie(w, &cookie)
	http.ServeFile(w, r, "index.html")
}
