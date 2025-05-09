package user

import (
	"net/http"

	"github.com/Code-Linx/Go-Ecommerce-Backend-Api/types"
	"github.com/Code-Linx/Go-Ecommerce-Backend-Api/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {}


func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	//get Json Payload
	var payload types.RegisterUserPayload

	if err := utils.ParseJson(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	//check if user exist
}
