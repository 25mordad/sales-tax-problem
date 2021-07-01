package receipt

import (
	"math"
)

//ProductType is an enum
type ProductType string

//We can add more product type here:
const (
	Book    ProductType = "Book"
	Food                = "Food"
	Medical             = "Medical"
	Other               = "Other"
)

// Product is for input struct
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

// ProductReceipt is a product struct uses for output inside the Receipts
type ProductReceipt struct {
	quantity int64
	name     string
	price    float64
}

//Receipt is our output
type Receipt struct {
	products []ProductReceipt
	tax      float64
	total    float64
}

//GetReceipt input is an array of input products and it return a type of Receipt
func GetReceipt(products []Product) Receipt {
	var receipt Receipt
	var receiptProduct ProductReceipt
	var receiptProducts []ProductReceipt
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

//calculateTax update the tax and FinalPrice and taxRate of a product
func calculateTax(product *Product) {
	product.TaxRate = getTaxRate(product)
	product.Tax = ceil((float64(product.TaxRate)*product.Price)/100, 0.05)
	product.FinalPrice = round2Dec(product.Price + product.Tax)
}

//getTaxRate calculate the tax rate of a product based on our rules
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

//ceil calculate the nearest based on the unit (at this problem unit is 0.05)
func ceil(x, unit float64) float64 {
	return round2Dec((math.Ceil(x/unit) * unit))
}

//round2Dec to round a number to 2 decimal digits
func round2Dec(x float64) float64 {
	return math.Round(x*100) / 100
}
