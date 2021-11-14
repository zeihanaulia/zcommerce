package order

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Order struct {
	host string
}

func NewOrder(host string) *Order {
	return &Order{host: host}
}

type OrderPlacedRequest struct {
	PaymentTrxID string `json:"payment_trx_id"`
}

type OrderPlacedResponse struct {
	Data struct {
		PaymentTrxID string `json:"payment_trx_id"`
	} `json:"data"`
}

func (o *Order) Placed(ctx context.Context, paymentTrxID string) error {
	postBody, _ := json.Marshal(OrderPlacedRequest{PaymentTrxID: paymentTrxID})
	responseBody := bytes.NewBuffer(postBody)
	url := fmt.Sprintf("%s/order/placed", o.host)
	fmt.Println(url)
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("POST", url, responseBody)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.ereg1pNERs1tmvgEBK9OX0-pCCnGqOKHlk7b0fUDUc8")
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	defer resp.Body.Close()

	var res OrderPlacedResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("payment.order.placed")

	return nil
}
