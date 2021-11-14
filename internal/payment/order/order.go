package order

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
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
