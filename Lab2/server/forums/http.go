package forums

import (
	"log"
	"net/http"
)

func HttpHandler(storage *Storage) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListForums(r, rw, storage)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}


func handleListForums(r *http.Request, rw http.ResponseWriter, storage *Storage) {

	res, err := storage.ListForums()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}
