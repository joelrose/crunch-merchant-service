package menus

import (
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
	"github.com/labstack/gommon/log"
)

func (s *MenuService) getChildrenRecursive(productMap map[uuid.UUID]dtos.GetStoreProduct, productId uuid.UUID) dtos.GetStoreProduct {
	childrenIds, err := s.db.GetProductChildren(productId)
	if err != nil {
		log.Errorf("failed to get product children: %v", err)
		return dtos.GetStoreProduct{}
	}

	product := productMap[productId]
	result := []dtos.GetStoreProduct{}
	for _, childId := range childrenIds {
		result = append(result, s.getChildrenRecursive(productMap, childId))
	}

	if len(result) > 0 {
		product.Products = result
	}

	return product
}

func (s *MenuService) build() (*MenuRedisModel, error) {
	log.Debugf("rebuilding menu: [%v]", s.storeId)

	categories, err := s.db.GetCategories(s.storeId)
	if err != nil {
		log.Errorf("failed to get categories: %v", err)
		return nil, err
	}

	products, err := s.db.GetProducts(s.storeId)
	if err != nil {
		log.Errorf("failed to get products: %v", err)
		return nil, err
	}

	productMap := make(map[uuid.UUID]dtos.GetStoreProduct, len(products))
	for _, product := range products {
		productMap[product.Id] = product
	}

	for ind, category := range categories {
		childrenProductIds, err := s.db.GetCategoryChildren(category.Id)

		if err != nil {
			log.Errorf("failed to get category children: %v", err)
			return nil, err
		}

		var children []dtos.GetStoreProduct
		for _, childrenProductId := range childrenProductIds {
			children = append(children, s.getChildrenRecursive(productMap, childrenProductId))
		}

		if len(children) > 0 {
			categories[ind].Products = children
		}
	}

	openingHours, err := s.db.GetOpeningHours(s.storeId)
	if err != nil {
		log.Errorf("failed to get opening hours: %v", err)
		return nil, err
	}

	return &MenuRedisModel{
		Categories:   categories,
		OpeningHours: openingHours,
	}, nil
}
