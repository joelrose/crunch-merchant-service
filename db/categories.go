package db

import "github.com/joelrose/crunch-merchant-service/db/models"

func (db *DB) CreateCategory(category models.MenuCategory) (int, error) {
	var lastInsertId int
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

func (db *DB) DeleteCategories(storeId int) error {
	_, err := db.Sqlx.Exec("DELETE FROM menu_categories WHERE store_id = $1", storeId)

	return err
}

func (db *DB) CreateProductCategoryRelation(categoryId int, productId int) error {
	_, err := db.Sqlx.Exec(
		"INSERT INTO category_product_relation (menu_category_id, menu_product_id) VALUES ($1, $2)",
		categoryId, productId,
	)

	return err
}

func (db *DB) GetCategories(storeId int) ([]models.MenuCategory, error) {
	var categories []models.MenuCategory
	err := db.Sqlx.Select(&categories, "SELECT * FROM menu_categories WHERE store_id = $1", storeId)

	return categories, err
}

func (db *DB) GetCategoryChildren(categoryId int) ([]string, error) {
	var categoryRelations []string
	err := db.Sqlx.Select(&categoryRelations, "SELECT menu_product_id FROM category_product_relation WHERE menu_category_id = $1", categoryId)

	return categoryRelations, err
}
