package users

import (
	"encoding/json"
	"log"
	"net/http"
)

func HttpHandler(storage *Storage) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handleCreateUser(r, rw, storage)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}


func handleCreateUser(r *http.Request, rw http.ResponseWriter, storage *Storage) {

	var u User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.Printf("Error decoding user input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}

	err := storage.CreateUser(u.Name, u.Interests)

	if err == nil {
		tools.WriteJsonOk(rw, &u)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}
