package service

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/zeihanaulia/zcommerce/internal/account"
)

type AccountRepository interface {
	Create(ctx context.Context, accounts account.Account) (account.Account, error)
	FindByEmail(ctx context.Context, email string) (account.Account, error)
}

type Account struct {
	account AccountRepository
}

func NewAccount(account AccountRepository) *Account {
	return &Account{account: account}
}

func (a *Account) Register(ctx context.Context, accounts account.Account) (account.Account, error) {
	resp, err := a.account.Create(ctx, accounts)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (a *Account) Login(ctx context.Context, email, password string) (string, error) {
	accounts, err := a.account.FindByEmail(ctx, email)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if accounts.Password == password {
		token, err := CreateToken(email)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		return token, nil
	}

	return "", nil
}

func CreateToken(email string) (string, error) {
	var err error

	var secret = "rahasia"
	atClaims := jwt.MapClaims{}
	// atClaims["authorized"] = true
	// atClaims["user_id"] = email
	// atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}
