package db

import (
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/models"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
)

type DBInterface interface {
	// Categories
	CreateCategory(category models.MenuCategory) (uuid.UUID, error)
	DeleteCategories(storeId uuid.UUID) error
	CreateProductCategoryRelation(categoryId uuid.UUID, productId uuid.UUID) error
	GetCategories(storeId uuid.UUID) ([]dtos.GetStoreCategory, error)
	GetCategoryChildren(categoryId uuid.UUID) ([]uuid.UUID, error)

	// Channels
	GetChannelByStoreId(storeId uuid.UUID) (models.DeliverectChannel, error)
	GetChannelByDeliverectLinkId(deliverectLinkId string) (models.DeliverectChannel, error)
	CreateChannel(storeId uuid.UUID, locationId string, deliverectChannelLinkId string, status dtos.ChannelStatus) error
	UpdateChannelStatus(status dtos.ChannelStatus, storeId uuid.UUID) error

	// Products
	CreateProduct(product models.MenuProduct) (uuid.UUID, error)
	DeleteProducts(storeId uuid.UUID) error
	CreateProductRelation(childProductId uuid.UUID, parentProductId uuid.UUID) error
	GetProducts(storeId uuid.UUID) ([]dtos.GetStoreProduct, error)
	GetTopProducts(storeId uuid.UUID) ([]dtos.GetStoreProduct, error)
	GetProductsByPlu(plu string, storeId uuid.UUID) ([]uuid.UUID, error)
	GetProductChildren(parentProductId uuid.UUID) ([]uuid.UUID, error)
	UpdateProductsSnooze(productIds []uuid.UUID, snooze bool) error
	UpdateProductSortOrder(childProductId uuid.UUID, sortOrder int) error

	// Order items
	CreateOrderItemWithoutParent(orderItem models.CreateOrderItem) (uuid.UUID, error)
	CreateOrderItemWithParent(orderItem models.CreateOrderItem) (uuid.UUID, error)
	GetOrderItems(orderId uuid.UUID) ([]models.OrderItem, error)

	// Orders
	CreateOrder(order models.CreateOrder) (uuid.UUID, error)
	GetOrderByStripeOrderId(stripeOrderId string) (models.Order, error)
	GetOrdersByUserId(userId uuid.UUID) ([]dtos.GetOrdersResponse, error)
	GetOrdersByStoreId(storeId uuid.UUID) ([]dtos.GetOrdersResponse, error)
	GetOrderById(orderId uuid.UUID) (models.Order, error)
	UpdateOrderStatus(orderId uuid.UUID, orderStatus models.OrderStatus) error
	MarkOrderAsPaid(orderId uuid.UUID) error

	// Store Opening Hours
	CreateStoreOpeningHour(openingHour models.StoreOpeningHour) error
	DeleteOpeningHours(storeId uuid.UUID) error
	GetOpeningHours(storeId uuid.UUID) ([]dtos.GetStoreOpeningHour, error)

	// Stores
	GetStoreById(storeId uuid.UUID) (models.Store, error)
	GetStoreByMerchantUserId(merchantUserId string) (uuid.UUID, error)
	GetOpenStore(storeId uuid.UUID) (models.Store, error)
	GetAvailableStore(storeId uuid.UUID) (models.Store, error)
	GetOpenStores() ([]dtos.GetStoresOverviewResponse, error)
	SetIsOpen(storeId uuid.UUID, isOpen bool) error
	SetStoreImageUrl(storeId uuid.UUID, imageUrl string) error

	// Users
	GetAllUsers() ([]models.User, error)
	GetUserByFirebaseId(firebaseId string) (models.User, error)
	GetUserByUserId(userId uuid.UUID) (models.User, error)
	CreateUser(firebaseId string, user dtos.CreateUserRequest) error

	// Whitelist
	IsWhitelisted(identifier string) bool
}
