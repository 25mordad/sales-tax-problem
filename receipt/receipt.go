package receipt

type ProductType string

const (
	Book    ProductType = "Book"
	Food                = "Food"
	Medical             = "Medical"
	Other               = "Other"
)

type Product struct {
	Quantity   int64
	Name       string
	Ptype      ProductType
	Price      float64
	IsImported bool
	TaxRate    int64
	Tax        float64
	FinalPrice float64
}

type ReceiptProduct struct {
	quantity int64
	name     string
	price    float64
}

type Receipt struct {
	products []ReceiptProduct
	tax      float64
	total    float64
}

func GetReceipt(products []Product) Receipt {
	var receipt Receipt

	return receipt
}
