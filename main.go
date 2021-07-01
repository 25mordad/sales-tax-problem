package main

import (
	"fmt"
	rcp "sales-tax/receipt"
)

///
func main() {

	var p1 rcp.Product
	p1.Quantity = 1
	p1.Name = "bottle of perfume"
	p1.Ptype = rcp.Other
	p1.Price = 27.99
	p1.IsImported = true
	fmt.Println("p1", p1)

	var p2 rcp.Product
	p2.Quantity = 1
	p2.Name = "perfume"
	p2.Ptype = rcp.Other
	p2.Price = 18.99
	p2.IsImported = false
	fmt.Println("p2", p2)

	var p3 rcp.Product
	p3.Quantity = 1
	p3.Name = "packet of headache pills"
	p3.Ptype = rcp.Medical
	p3.Price = 9.75
	p3.IsImported = false
	fmt.Println("p3", p3)

	var p4 rcp.Product
	p4.Quantity = 1
	p4.Name = "box of imported chocolates"
	p4.Ptype = rcp.Food
	p4.Price = 11.25
	p4.IsImported = true
	fmt.Println("p4", p4)

	var products []rcp.Product
	products = append(products, p1)
	products = append(products, p2)
	products = append(products, p3)
	products = append(products, p4)
	fmt.Printf("products %+v \n\n", products)
	fmt.Printf("getReceipt %+v \n\n", rcp.GetReceipt(products))

}
