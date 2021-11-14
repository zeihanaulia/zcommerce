package postgresql

import (
	"context"

	"github.com/zeihanaulia/zcommerce/internal/account"
	"github.com/zeihanaulia/zcommerce/internal/account/postgresql/db"
)

// Account represent orders tables
type Account struct {
	q *db.Queries
}

func NewAccount(d db.DBTX) *Account {
	return &Account{
		q: db.New(d),
	}
}

func (a *Account) Create(ctx context.Context, accounts account.Account) (account.Account, error) {
	_, err := a.q.CreateAccounts(ctx, db.CreateAccountsParams{
		Email:    accounts.Email,
		Password: accounts.Password,
	})
	return account.Account{
		Email: accounts.Email,
	}, err
}

func (a *Account) FindByEmail(ctx context.Context, email string) (account.Account, error) {
	res, err := a.q.SelectAccounts(ctx, email)
	if err != nil {
		return account.Account{}, nil
	}

	return account.Account{
		Email:    res.Email,
		Password: res.Password,
	}, err
}
