package product

import (
	"fmt"
)

type Service struct {
}

func NewService() *Service {

	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}
func (s *Service) GetProductTitle(id int) (string, error) {
	if id < 0 || id > len(allProducts) {
		return "", fmt.Errorf("wrong product id: %d", id)
	}
	return allProducts[id-1].Title, nil
}
