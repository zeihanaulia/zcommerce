// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"github.com/jackc/pgtype"
)

type Order struct {
	ID              int32       `json:"id"`
	TrxID           string      `json:"trx_id"`
	PaymentTrxID    string      `json:"payment_trx_id"`
	LockItems       pgtype.JSON `json:"lock_items"`
	Status          string      `json:"status"`
	CustomerName    string      `json:"customer_name"`
	CustomerAddress string      `json:"customer_address"`
}