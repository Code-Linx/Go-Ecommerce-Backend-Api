package cart

import (
	"fmt"
	"net/http"

	"github.com/Code-Linx/Go-Ecommerce-Backend-Api/types"
	"github.com/Code-Linx/Go-Ecommerce-Backend-Api/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
	store        types.OrderStore
	productStore types.ProductStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore) *Handler {
	return &Handler{store: store, productStore: productStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", h.handleCheckOut).Methods(http.MethodPost)
}

func (h *Handler) handleCheckOut(w http.ResponseWriter, r *http.Request) {
	var cart types.CartCheckoutPayload
	if err := utils.ParseJson(r, cart); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload : %v", errors))
		return
	}

	//get product
	/* productIDs, err := getCartItemsIDs(cart.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	} */

}
