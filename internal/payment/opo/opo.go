package opo

import (
	"context"
	"fmt"
	"time"
)

type OPO struct {
}

func NewOPO() *OPO {
	return &OPO{}
}

func (o *OPO) Paid(ctx context.Context, phoneNumber string) (string, error) {
	now := time.Now()
	trxID := fmt.Sprintf("OPO-%d", now.Unix()) // simulating opo success
	return trxID, nil
}
