package product

import (
	"github.com/funmi4194/product-catalog/models"
	"github.com/funmi4194/product-catalog/sorter"
)

// IProductService is an interface that defines product related logic asides from sorting
type IProductService interface {
	SortProducts(products []models.Product, strategy sorter.SortType) []models.Product
}

type ProductService struct{}

// NewProductService returns a new instance of ProductService
func NewProductService() IProductService {
	return &ProductService{}
}

// SortProducts is a product logic delegated to sorting tasks to a given sort type
func (s *ProductService) SortProducts(products []models.Product, strategy sorter.SortType) []models.Product {
	return strategy.Sort(products)
}
