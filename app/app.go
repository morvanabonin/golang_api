package app

import "net/http"

// add = post
// update = put
// list = get

func ListHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Pokemon!\n"))
	
	if err != nil {
		return
	}
}
