package menus

import (
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/dtos"
	"github.com/labstack/gommon/log"
)

func getChildrenRecursive(db *db.DB, productMap map[uuid.UUID]dtos.GetStoreProduct, productId uuid.UUID) dtos.GetStoreProduct {
	childrenIds, err := db.GetProductChildren(productId)
	if err != nil {
		log.Errorf("failed to get product children: %v", err)
		return dtos.GetStoreProduct{}
	}

	product := productMap[productId]
	result := []dtos.GetStoreProduct{}
	for _, childId := range childrenIds {
		result = append(result, getChildrenRecursive(db, productMap, childId))
	}

	if len(result) > 0 {
		product.Products = result
	}

	return product
}

func Build(db *db.DB, storeId uuid.UUID) (*MenuRedisModel, error) {
	categories, err := db.GetCategories(storeId)
	if err != nil {
		log.Errorf("failed to get categories: %v", err)
		return nil, err
	}

	products, err := db.GetProducts(storeId)
	if err != nil {
		log.Errorf("failed to get products: %v", err)
		return nil, err
	}

	productMap := map[uuid.UUID]dtos.GetStoreProduct{}
	for _, product := range products {
		productMap[product.Id] = product
	}

	for ind, category := range categories {
		childrenProductIds, err := db.GetCategoryChildren(category.Id)

		if err != nil {
			log.Errorf("failed to get category children: %v", err)
			return nil, err
		}

		children := []dtos.GetStoreProduct{}
		for _, childrenProductId := range childrenProductIds {
			children = append(children, getChildrenRecursive(db, productMap, childrenProductId))
		}

		if len(children) > 0 {
			categories[ind].Products = children
		}
	}

	openingHours, err := db.GetOpeningHours(storeId)
	if err != nil {
		log.Errorf("failed to get opening hours: %v", err)
		return nil, err
	}

	return &MenuRedisModel{
		Categories:   categories,
		OpeningHours: openingHours,
	}, nil
}
