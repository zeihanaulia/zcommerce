/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"
	"github.com/zeihanaulia/zcommerce/adapter"
	"github.com/zeihanaulia/zcommerce/internal/order/payment"
	"github.com/zeihanaulia/zcommerce/internal/order/postgresql"
	"github.com/zeihanaulia/zcommerce/internal/order/rest"
	"github.com/zeihanaulia/zcommerce/internal/order/service"
)

// orderRestCmd represents the orderRest command
var orderRestCmd = &cobra.Command{
	Use:   "order-rest",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		r := chi.NewRouter()

		// Infrastructure
		db, err := adapter.NewPostgreSQL()
		if err != nil {
			log.Println(err)
		}

		// Repository
		orders := postgresql.NewOrder(db)
		payments := payment.NewPayment("http://localhost:8003")

		// Service
		scv := service.NewOrder(orders, payments)

		// Handler
		handler := rest.NewOrderHandler(scv)
		handler.Register(r)

		// Healthz
		r.Get("/healthz", func(rw http.ResponseWriter, r *http.Request) {
			// TODO: check all infrastructure health
			// https://microservices.io/patterns/observability/health-check-api.html
			// https://openliberty.io/docs/21.0.0.10/health-check-microservices.html
			// https://github.com/etherlabsio/healthcheck
			rw.Write([]byte("ok!"))
		})

		// Metrics
		r.Get("/metrics", func(rw http.ResponseWriter, r *http.Request) {
			// TODO: collect metrics wiht prometheus
			// https://www.weave.works/blog/the-red-method-key-metrics-for-microservices-architecture/
			rw.Write([]byte("ok!"))
		})

		// Server
		address := "0.0.0.0:8002"
		srv := &http.Server{
			Handler:           r,
			Addr:              address,
			ReadTimeout:       1 * time.Second,
			ReadHeaderTimeout: 1 * time.Second,
			WriteTimeout:      1 * time.Second,
			IdleTimeout:       1 * time.Second,
		}

		log.Printf("Listening on %s \n", address)
		log.Fatal(srv.ListenAndServe())
	},
}

func init() {
	rootCmd.AddCommand(orderRestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// orderRestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// orderRestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
