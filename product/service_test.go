package product

import (
	"testing"

	"github.com/funmi4194/product-catalog/models"
	"github.com/funmi4194/product-catalog/sorter"
	"github.com/stretchr/testify/assert"
)

func sampleProducts() []models.Product {
	return []models.Product{
		{ID: 1, Name: "Alabaster", Price: 20.00, SalesCount: 100, ViewsCount: 1000}, // 0.10
		{ID: 2, Name: "Zebra", Price: 10.00, SalesCount: 300, ViewsCount: 1000},     // 0.30
		{ID: 3, Name: "Coffee", Price: 30.00, SalesCount: 200, ViewsCount: 1000},    // 0.20
	}
}

func TestProductService_SortProducts(t *testing.T) {
	type testCase struct {
		name        string
		strategy    sorter.SortType
		expectedIDs []int
	}

	cases := []testCase{
		{
			name:        "sort by price ascending",
			strategy:    sorter.ByPrice{},
			expectedIDs: []int{2, 1, 3},
		},
		{
			name:        "sort by conversion descending",
			strategy:    sorter.ByConversion{},
			expectedIDs: []int{2, 3, 1},
		},
	}

	service := NewProductService()

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			sorted := service.SortProducts(sampleProducts(), tc.strategy)
			var gotIDs []int
			for _, p := range sorted {
				gotIDs = append(gotIDs, p.ID)
			}
			assert.Equal(t, tc.expectedIDs, gotIDs)
		})
	}
}

func TestProductService_EmptyInput(t *testing.T) {
	service := NewProductService()
	sorted := service.SortProducts([]models.Product{}, sorter.ByPrice{})
	assert.Empty(t, sorted)
}
