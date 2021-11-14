package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zeihanaulia/zcommerce/internal/account"
)

type AccountService interface {
	Register(ctx context.Context, accounts account.Account) (account.Account, error)
	Login(ctx context.Context, email, password string) (string, error)
}

type AccountHandler struct {
	svc AccountService
}

func NewAccountHandler(svc AccountService) *AccountHandler {
	return &AccountHandler{svc: svc}
}

func (a *AccountHandler) Register(r chi.Router) {
	r.Route("/account", func(r chi.Router) {
		r.Post("/register", a.register)
		r.Post("/login", a.login)
	})
}

type AccountRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Data struct {
		ID    int64  `json:"id"`
		Email string `json:"email"`
	} `json:"data"`
}

func (a *AccountHandler) register(w http.ResponseWriter, r *http.Request) {
	var req AccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("decoder.order.placed: ", err)
		return
	}

	fmt.Println("account.register")

	resp, err := a.svc.Register(r.Context(), account.Account{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		log.Println(err)
		return
	}

	renderResponse(w, RegisterResponse{
		Data: struct {
			ID    int64  "json:\"id\""
			Email string "json:\"email\""
		}{
			ID:    int64(resp.ID),
			Email: resp.Email,
		},
	}, http.StatusOK)
}

type LoginResponse struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

func (a *AccountHandler) login(w http.ResponseWriter, r *http.Request) {
	var req AccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("decoder.order.placed: ", err)
		return
	}

	resp, err := a.svc.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		log.Println(err)
		return
	}

	renderResponse(w, LoginResponse{
		Data: struct {
			Token string "json:\"token\""
		}{
			Token: resp,
		},
	}, http.StatusOK)
}
