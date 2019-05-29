package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"simple-go-restapi/app/repository"
	userRepo "simple-go-restapi/app/repository/userrepository"
	"simple-go-restapi/app/sqldriver"
)

func NewHandler(db *sqldriver.DB) *Handler {
	return &Handler{
		repo: userRepo.NewSQLRepo(db.SQL),
	}
}

type Handler struct {
	repo repository.UserRepo
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	payload, err := h.repo.GetByName(r.Context(), vars["name"])

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")

	} else {
		respondWithJSON(w, http.StatusOK, payload)
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"message": msg})
}

