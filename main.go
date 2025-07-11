package main

import (
	"fmt"

	"github.com/funmi4194/product-catalog/models"
	"github.com/funmi4194/product-catalog/product"
	"github.com/funmi4194/product-catalog/sorter"
)

func main() {
	// define product to sort
	products := []models.Product{
		{ID: 1, Name: "Alabaster Table", Price: 12.99, Created: "2019-01-04", SalesCount: 32, ViewsCount: 730},
		{ID: 2, Name: "Zebra Table", Price: 44.49, Created: "2012-01-04", SalesCount: 301, ViewsCount: 3279},
		{ID: 3, Name: "Coffee Table", Price: 10.00, Created: "2014-05-28", SalesCount: 1048, ViewsCount: 20123},
	}

	service := product.NewProductService()

	sortType := sorter.ByPrice{}
	// sortType := sorter.ByConversion{}

	output := service.SortProducts(products, sortType)

	fmt.Println(output)
}
