package user

import (
	"fmt"
	"net/http"

	"github.com/golang_api/Ecommerce/types"
	"github.com/golang_api/Ecommerce/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/user", h.handleUser).Methods("GET")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleUser(w http.ResponseWriter, r *http.Request) {
	var user *[]types.User
	//user.ID = 12
	user, err := h.store.GetAllUser()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	err = utils.WriteJSON(w, http.StatusAccepted, user)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payLoad types.RegisterUserPayload

	if err := utils.ParseJOSN(r, payLoad); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	_, err := h.store.GetUserByUserId(payLoad.ID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("no such user found"))
	}

	// err = h.store.CreateUser(types.User{
	// 	USERNAME:    u.USERNAME,
	// 	USER_ID:     u.USER_ID,
	// 	INSERTED_ON: u.INSERTED_ON,
	// 	UPDATED_ON:  u.UPDATED_ON,
	// })
}
