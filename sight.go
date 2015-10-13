package main

import (
	//"fmt"
	"net/http"
)

func init() {
	myHandlers["/sight"] = func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		data := r.Form
		file_content := data.Get("file")

		w.Write([]byte(file_content))
	}
}
