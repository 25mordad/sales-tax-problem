package receipt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReceipt(t *testing.T) {
	assert := assert.New(t)

	var p1 Product
	p1.Quantity = 1
	p1.Name = "bottle of perfume"
	p1.Ptype = Other
	p1.Price = 27.99
	p1.IsImported = true

	var p2 Product
	p2.Quantity = 1
	p2.Name = "perfume"
	p2.Ptype = Other
	p2.Price = 18.99
	p2.IsImported = false

	var p3 Product
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

	answ := GetReceipt(products)
	fmt.Printf("getReceipt %+v \n\n", answ)
	assert.Equal(int64(1), answ.products[0].quantity)
	assert.Equal("Imported bottle of perfume", answ.products[0].name)
	assert.Equal(32.19, answ.products[0].price)
	assert.Equal(20.89, answ.products[1].price)
	assert.Equal(9.75, answ.products[2].price)
	assert.Equal(11.85, answ.products[3].price)
	assert.Equal(6.7, answ.tax)
	assert.Equal(74.68, answ.total)

}
