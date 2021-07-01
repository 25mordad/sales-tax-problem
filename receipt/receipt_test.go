package receipt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReceipt(t *testing.T) {
	assert := assert.New(t)

	////test input 1
	var p1 Product
	p1.Quantity = 1
	p1.Name = "book"
	p1.Ptype = Book
	p1.Price = 12.49
	p1.IsImported = false

	var p2 Product
	p2.Quantity = 1
	p2.Name = "music CD"
	p2.Ptype = Other
	p2.Price = 14.99
	p2.IsImported = false

	var p3 Product
	p3.Quantity = 1
	p3.Name = "chocolate bar"
	p3.Ptype = Food
	p3.Price = 0.85
	p3.IsImported = false

	var products1 []Product
	products1 = append(products1, p1)
	products1 = append(products1, p2)
	products1 = append(products1, p3)

	answ := GetReceipt(products1)
	fmt.Printf("getReceipt %+v \n\n", answ)
	assert.Equal(int64(1), answ.products[0].quantity)
	assert.Equal("book", answ.products[0].name)
	assert.Equal(12.49, answ.products[0].price)
	assert.Equal(16.49, answ.products[1].price)
	assert.Equal(0.85, answ.products[2].price)
	assert.Equal(1.5, answ.tax)
	assert.Equal(29.83, answ.total)

	/////input 2
	p1.Quantity = 1
	p1.Name = "box of chocolates"
	p1.Ptype = Food
	p1.Price = 10.00
	p1.IsImported = true

	p2.Quantity = 1
	p2.Name = "bottle of perfume"
	p2.Ptype = Other
	p2.Price = 47.5
	p2.IsImported = true

	var products2 []Product
	products2 = append(products2, p1)
	products2 = append(products2, p2)

	answ = GetReceipt(products2)
	fmt.Printf("getReceipt %+v \n\n", answ)
	assert.Equal(int64(1), answ.products[0].quantity)
	assert.Equal("Imported box of chocolates", answ.products[0].name)
	assert.Equal(10.5, answ.products[0].price)
	assert.Equal(54.65, answ.products[1].price)
	assert.Equal(7.65, answ.tax)
	assert.Equal(65.15, answ.total)

	////// input 3
	p1.Quantity = 1
	p1.Name = "bottle of perfume"
	p1.Ptype = Other
	p1.Price = 27.99
	p1.IsImported = true

	p2.Quantity = 1
	p2.Name = "perfume"
	p2.Ptype = Other
	p2.Price = 18.99
	p2.IsImported = false

	p3.Quantity = 1
	p3.Name = "packet of headache pills"
	p3.Ptype = Medical
	p3.Price = 9.75
	p3.IsImported = false

	var p4 Product
	p4.Quantity = 1
	p4.Name = "box of imported chocolates"
	p4.Ptype = Food
	p4.Price = 11.25
	p4.IsImported = true

	var products []Product
	products = append(products, p1)
	products = append(products, p2)
	products = append(products, p3)
	products = append(products, p4)

	answ = GetReceipt(products)
	fmt.Printf("getReceipt %+v \n\n", answ)
	assert.Equal(int64(1), answ.products[0].quantity)
	assert.Equal("Imported bottle of perfume", answ.products[0].name)
	assert.Equal(32.19, answ.products[0].price)
	assert.Equal(20.89, answ.products[1].price)
	assert.Equal(9.75, answ.products[2].price)
	assert.Equal(11.85, answ.products[3].price)
	assert.Equal(6.7, answ.tax)
	assert.Equal(74.68, answ.total)

	/////input 4: test quantity
	p1.Quantity = 2
	p1.Name = "box of chocolates"
	p1.Ptype = Food
	p1.Price = 10.00
	p1.IsImported = true

	p2.Quantity = 1
	p2.Name = "bottle of perfume"
	p2.Ptype = Other
	p2.Price = 47.5
	p2.IsImported = true

	var products3 []Product
	products3 = append(products3, p1)
	products3 = append(products3, p2)

	answ = GetReceipt(products3)
	fmt.Printf("getReceipt %+v \n\n", answ)
	assert.Equal(int64(2), answ.products[0].quantity)
	assert.Equal("Imported box of chocolates", answ.products[0].name)
	assert.Equal(21.0, answ.products[0].price)
	assert.Equal(54.65, answ.products[1].price)
	assert.Equal(8.15, answ.tax)
	assert.Equal(75.65, answ.total)

}
