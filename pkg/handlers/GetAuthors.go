package handlers

import (
	"ebooks/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h handler) GetAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author

	if result := h.DB.Find(&authors); result.Error != nil {
		// don't use print lines except for debugging
		// you don't want to just continue here either, handle this error.
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// error handeling
	json.NewEncoder(w).Encode(authors)
}
