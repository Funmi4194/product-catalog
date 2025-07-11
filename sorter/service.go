package sorter

import (
	"sort"

	"github.com/funmi4194/product-catalog/models"
)

type SortType interface {
	Sort(products []models.Product) []models.Product
}

// ByConversion sorts products by sales/view ratio in descending order.
type ByConversion struct{}

func (s ByConversion) Sort(products []models.Product) []models.Product {
	sort.Slice(products, func(i, j int) bool {
		ratioA := float64(products[i].SalesCount) / float64(products[i].ViewsCount)
		ratioB := float64(products[j].SalesCount) / float64(products[j].ViewsCount)
		return ratioA > ratioB
	})
	return products
}

// ByPrice sorts products by ascending price.
type ByPrice struct{}

func (s ByPrice) Sort(products []models.Product) []models.Product {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})
	return products
}

// more sort technic can be added without touching the previously implemented sort techinics
