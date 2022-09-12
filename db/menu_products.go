package db

import (
	"github.com/joelrose/crunch-merchant-service/db/models"
)

func (db *DB) CreateProduct(product models.MenuProduct) (int, error) {
	var lastInsertId int
	err := db.Sqlx.Get(
		&lastInsertId,
		"INSERT INTO menu_product (name, plu, price, description, snoozed, tax, image_url, max, min, multiply, product_type, sort_order, visible, store_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id",
		product.Name,
		product.Plu,
		product.Price,
		product.Description,
		product.Snoozed,
		product.Tax,
		product.ImageUrl,
		product.Max,
		product.Min,
		product.Multiply,
		product.ProductType,
		product.SortOrder,
		product.Visible,
		product.StoreId,
	)

	return lastInsertId, err
}

func (db *DB) DeleteProducts(storeId int) error {
	_, err := db.Sqlx.Exec("DELETE FROM menu_product WHERE store_id = $1", storeId)

	return err
}

func (db *DB) CreateProductRelation(childProductId int, parentProductId int) error {
	_, err := db.Sqlx.Exec(
		"INSERT INTO product_product_relation (child_product_id, parent_product_id) VALUES ($1, $2)",
		childProductId, parentProductId,
	)

	return err
}
