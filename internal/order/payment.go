package order

// Payment defines a type of payment service
// The service should to know what customer order
// and Locking items before payment
type Payment struct {
	TransactionDetail TransactionDetail
	CustomerDetail    CustomerDetail
	ItemDetails       []ItemDetail
}

// TransactionDetail defines type of transaction detail
// such as transaction id as a key
type TransactionDetail struct {
	TrxID       string
	FinalAmount float64
}

// CustomerDetail defines type of customer detail
type CustomerDetail struct {
	ID      string
	Name    string
	Address string
}

// ItemDetail defines type of items detail
type ItemDetail struct {
	ID       string
	Name     string
	Quantity float64
	Uom      string
	Price    float64
}
