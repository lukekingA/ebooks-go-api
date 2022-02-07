package handlers

import (
	"ebooks/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) GetAuthorById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// error handeling
	id, _ := strconv.Atoi(vars["id"])

	var author models.Author

	if result := h.DB.First(&author, id); result.Error != nil {
		// error handeling
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// error handeling
	json.NewEncoder(w).Encode(author)
}
