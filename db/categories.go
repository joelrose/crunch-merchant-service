package db

import (
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/models"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
)

func (db *DB) CreateCategory(category models.MenuCategory) (uuid.UUID, error) {
	var lastInsertId uuid.UUID
	err := db.Sqlx.Get(
		&lastInsertId,
		"INSERT INTO menu_categories (name, description, image_url, sort_order, store_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		category.Name,
		category.Description,
		category.ImageUrl,
		category.SortOrder,
		category.StoreId,
	)

	return lastInsertId, err
}

func (db *DB) DeleteCategories(storeId uuid.UUID) error {
	_, err := db.Sqlx.Exec("DELETE FROM menu_categories WHERE store_id = $1", storeId)

	return err
}

func (db *DB) CreateProductCategoryRelation(categoryId uuid.UUID, productId uuid.UUID) error {
	_, err := db.Sqlx.Exec(
		"INSERT INTO category_product_relation (menu_category_id, menu_product_id) VALUES ($1, $2)",
		categoryId, productId,
	)

	return err
}

func (db *DB) GetCategories(storeId uuid.UUID) ([]dtos.GetStoreCategory, error) {
	var categories []dtos.GetStoreCategory
	err := db.Sqlx.Select(&categories, "SELECT id, name, description, image_url, sort_order FROM menu_categories WHERE store_id = $1", storeId)

	return categories, err
}

func (db *DB) GetCategoryChildren(categoryId uuid.UUID) ([]uuid.UUID, error) {
	var categoryRelations []uuid.UUID
	err := db.Sqlx.Select(&categoryRelations, "SELECT menu_product_id FROM category_product_relation WHERE menu_category_id = $1", categoryId)

	return categoryRelations, err
}
