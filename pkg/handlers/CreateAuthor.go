package handlers

import (
	"ebooks/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (h handler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	// This error should be cought and an approiate response returned to the user before the process terminated instad of just ditching here
	if err != nil {
		log.Fatalln(err)
	}

	var author models.Author
	// Json.unmarshal returnes an error. It's a rare case that you'll chose to ignore handling an error.
	json.Unmarshal(body, &author)

	// It would be nice if there were some kind of request validation here. What if the request contains an empty author. Probably Id and LastName atleast are required fields to Create this author record.

	if result := h.DB.Create(&author); result.Error != nil {
		// This if block would contain the failure response in case of db error response.
		fmt.Println(result.Error)
	}

	// send 201
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	// again don't ignore the error returned by Encode
	json.NewEncoder(w).Encode("Created")
}
