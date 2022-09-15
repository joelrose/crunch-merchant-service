package db

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/joelrose/crunch-merchant-service/db/models"
	"github.com/joelrose/crunch-merchant-service/dtos"
)

func (db *DB) CreateProduct(product models.MenuProduct) (int, error) {
	var lastInsertId int
	err := db.Sqlx.Get(
		&lastInsertId,
		"INSERT INTO menu_product (name, plu, price, description, snoozed, tax, image_url, max, min, multiply, multi_max, product_type, sort_order, visible, store_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING id",
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
		product.MultiMax,
		product.ProductType,
		product.SortOrder,
		product.Visible,
		product.StoreId,
	)

	return lastInsertId, err
}

func (db *DB) DeleteProducts(storeId uuid.UUID) error {
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

func (db *DB) GetProducts(storeId uuid.UUID) ([]dtos.GetStoreProduct, error) {
	var products []dtos.GetStoreProduct
	err := db.Sqlx.Select(
		&products,
		"SELECT id, name, description, price, max, min, multiply, multi_max, plu, snoozed, tax, product_type, image_url, sort_order,visible FROM menu_product WHERE store_id = $1",
		storeId,
	)

	return products, err
}

func (db *DB) GetProductsByPlu(plu string, storeId uuid.UUID) ([]int, error) {
	var productIds []int
	err := db.Sqlx.Select(&productIds, "SELECT id FROM menu_product WHERE plu LIKE $1 AND store_id = $2", "%"+plu+"%", storeId)

	return productIds, err
}

func (db *DB) GetProductChildren(parentProductId int) ([]int, error) {
	var productIds []int
	err := db.Sqlx.Select(&productIds, "SELECT child_product_id FROM product_product_relation WHERE parent_product_id = $1", parentProductId)

	return productIds, err
}

func (db *DB) UpdateProductsSnooze(productIds []int, snooze bool) error {
	query, args, err := sqlx.In("UPDATE menu_product SET snoozed = ? WHERE id IN (?)", snooze, productIds)
	if err != nil {
		return err
	}

	query = db.Sqlx.Rebind(query)
	_, err = db.Sqlx.Exec(query, args...)

	return err
}
