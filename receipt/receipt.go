package receipt

import (
	"math"
)

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
	var receiptProduct ReceiptProduct
	var receiptProducts []ReceiptProduct
	var salesTax float64
	var total float64
	salesTax = 0
	total = 0
	for _, product := range products {
		calculateTax(&product)
		receiptProduct.quantity = product.Quantity
		if product.IsImported {
			receiptProduct.name = "Imported " + product.Name
		} else {
			receiptProduct.name = product.Name
		}
		receiptProduct.price = product.FinalPrice * float64(product.Quantity)
		receiptProducts = append(receiptProducts, receiptProduct)
		salesTax = salesTax + product.Tax*float64(product.Quantity)
		total = total + product.FinalPrice*float64(product.Quantity)
	}
	receipt.products = receiptProducts
	receipt.tax = round2Dec(salesTax)
	receipt.total = round2Dec(total)

	return receipt
}

func calculateTax(product *Product) {
	product.TaxRate = getTaxRate(product)
	product.Tax = ceil((float64(product.TaxRate)*product.Price)/100, 0.05)
	product.FinalPrice = round2Dec(product.Price + product.Tax)
}

func getTaxRate(prd *Product) int64 {
	var taxRate int64
	taxRate = 0

	if prd.Ptype == Other {
		if prd.IsImported {
			taxRate = 15
		} else {
			taxRate = 10
		}

	} else {
		if prd.IsImported {
			taxRate = 5
		}
	}
	return taxRate
}

func ceil(x, unit float64) float64 {
	return round2Dec((math.Ceil(x/unit) * unit))
}

func round2Dec(x float64) float64 {
	return math.Round(x*100) / 100
}
