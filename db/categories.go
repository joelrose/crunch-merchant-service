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
