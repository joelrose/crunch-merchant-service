// Code generated by MockGen. DO NOT EDIT.
// Source: db/interface.go

// Package mock_db is a generated GoMock package.
package mock_db

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	models "github.com/joelrose/crunch-merchant-service/models"
	dtos "github.com/joelrose/crunch-merchant-service/models/dtos"
)

// MockDBInterface is a mock of DBInterface interface.
type MockDBInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDBInterfaceMockRecorder
}

// MockDBInterfaceMockRecorder is the mock recorder for MockDBInterface.
type MockDBInterfaceMockRecorder struct {
	mock *MockDBInterface
}

// NewMockDBInterface creates a new mock instance.
func NewMockDBInterface(ctrl *gomock.Controller) *MockDBInterface {
	mock := &MockDBInterface{ctrl: ctrl}
	mock.recorder = &MockDBInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBInterface) EXPECT() *MockDBInterfaceMockRecorder {
	return m.recorder
}

// CreateCategory mocks base method.
func (m *MockDBInterface) CreateCategory(category models.MenuCategory) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", category)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCategory indicates an expected call of CreateCategory.
func (mr *MockDBInterfaceMockRecorder) CreateCategory(category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockDBInterface)(nil).CreateCategory), category)
}

// CreateChannel mocks base method.
func (m *MockDBInterface) CreateChannel(storeId uuid.UUID, locationId, deliverectChannelLinkId string, status dtos.ChannelStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChannel", storeId, locationId, deliverectChannelLinkId, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateChannel indicates an expected call of CreateChannel.
func (mr *MockDBInterfaceMockRecorder) CreateChannel(storeId, locationId, deliverectChannelLinkId, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChannel", reflect.TypeOf((*MockDBInterface)(nil).CreateChannel), storeId, locationId, deliverectChannelLinkId, status)
}

// CreateOrder mocks base method.
func (m *MockDBInterface) CreateOrder(order models.CreateOrder) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", order)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockDBInterfaceMockRecorder) CreateOrder(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockDBInterface)(nil).CreateOrder), order)
}

// CreateOrderItemWithParent mocks base method.
func (m *MockDBInterface) CreateOrderItemWithParent(orderItem models.CreateOrderItem) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrderItemWithParent", orderItem)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrderItemWithParent indicates an expected call of CreateOrderItemWithParent.
func (mr *MockDBInterfaceMockRecorder) CreateOrderItemWithParent(orderItem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderItemWithParent", reflect.TypeOf((*MockDBInterface)(nil).CreateOrderItemWithParent), orderItem)
}

// CreateOrderItemWithoutParent mocks base method.
func (m *MockDBInterface) CreateOrderItemWithoutParent(orderItem models.CreateOrderItem) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrderItemWithoutParent", orderItem)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrderItemWithoutParent indicates an expected call of CreateOrderItemWithoutParent.
func (mr *MockDBInterfaceMockRecorder) CreateOrderItemWithoutParent(orderItem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderItemWithoutParent", reflect.TypeOf((*MockDBInterface)(nil).CreateOrderItemWithoutParent), orderItem)
}

// CreateProduct mocks base method.
func (m *MockDBInterface) CreateProduct(product models.MenuProduct) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", product)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockDBInterfaceMockRecorder) CreateProduct(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockDBInterface)(nil).CreateProduct), product)
}

// CreateProductCategoryRelation mocks base method.
func (m *MockDBInterface) CreateProductCategoryRelation(categoryId, productId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProductCategoryRelation", categoryId, productId)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProductCategoryRelation indicates an expected call of CreateProductCategoryRelation.
func (mr *MockDBInterfaceMockRecorder) CreateProductCategoryRelation(categoryId, productId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProductCategoryRelation", reflect.TypeOf((*MockDBInterface)(nil).CreateProductCategoryRelation), categoryId, productId)
}

// CreateProductRelation mocks base method.
func (m *MockDBInterface) CreateProductRelation(childProductId, parentProductId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProductRelation", childProductId, parentProductId)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProductRelation indicates an expected call of CreateProductRelation.
func (mr *MockDBInterfaceMockRecorder) CreateProductRelation(childProductId, parentProductId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProductRelation", reflect.TypeOf((*MockDBInterface)(nil).CreateProductRelation), childProductId, parentProductId)
}

// CreateStoreOpeningHour mocks base method.
func (m *MockDBInterface) CreateStoreOpeningHour(openingHour models.StoreOpeningHour) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStoreOpeningHour", openingHour)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateStoreOpeningHour indicates an expected call of CreateStoreOpeningHour.
func (mr *MockDBInterfaceMockRecorder) CreateStoreOpeningHour(openingHour interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStoreOpeningHour", reflect.TypeOf((*MockDBInterface)(nil).CreateStoreOpeningHour), openingHour)
}

// CreateUser mocks base method.
func (m *MockDBInterface) CreateUser(firebaseId string, user dtos.CreateUserRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", firebaseId, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockDBInterfaceMockRecorder) CreateUser(firebaseId, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDBInterface)(nil).CreateUser), firebaseId, user)
}

// DeleteCategories mocks base method.
func (m *MockDBInterface) DeleteCategories(storeId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategories", storeId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCategories indicates an expected call of DeleteCategories.
func (mr *MockDBInterfaceMockRecorder) DeleteCategories(storeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategories", reflect.TypeOf((*MockDBInterface)(nil).DeleteCategories), storeId)
}

// DeleteOpeningHours mocks base method.
func (m *MockDBInterface) DeleteOpeningHours(storeId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOpeningHours", storeId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOpeningHours indicates an expected call of DeleteOpeningHours.
func (mr *MockDBInterfaceMockRecorder) DeleteOpeningHours(storeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOpeningHours", reflect.TypeOf((*MockDBInterface)(nil).DeleteOpeningHours), storeId)
}

// DeleteProducts mocks base method.
func (m *MockDBInterface) DeleteProducts(storeId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProducts", storeId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProducts indicates an expected call of DeleteProducts.
func (mr *MockDBInterfaceMockRecorder) DeleteProducts(storeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProducts", reflect.TypeOf((*MockDBInterface)(nil).DeleteProducts), storeId)
}

// GetAllUsers mocks base method.
func (m *MockDBInterface) GetAllUsers() ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockDBInterfaceMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockDBInterface)(nil).GetAllUsers))
}

// GetAvailableStore mocks base method.
func (m *MockDBInterface) GetAvailableStore(storeId uuid.UUID) (models.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableStore", storeId)
	ret0, _ := ret[0].(models.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableStore indicates an expected call of GetAvailableStore.
func (mr *MockDBInterfaceMockRecorder) GetAvailableStore(storeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableStore", reflect.TypeOf((*MockDBInterface)(nil).GetAvailableStore), storeId)
}

// GetCategories mocks base method.
func (m *MockDBInterface) GetCategories(storeId uuid.UUID) ([]dtos.GetStoreCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategories", storeId)
	ret0, _ := ret[0].([]dtos.GetStoreCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories.
func (mr *MockDBInterfaceMockRecorder) GetCategories(storeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockDBInterface)(nil).GetCategories), storeId)
}

// GetCategoryChildren mocks base method.
func (m *MockDBInterface) GetCategoryChildren(categoryId uuid.UUID) ([]uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryChildren", categoryId)
	ret0, _ := ret[0].([]uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoryChildren indicates an expected call of GetCategoryChildren.
func (mr *MockDBInterfaceMockRecorder) GetCategoryChildren(categoryId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryChildren", reflect.TypeOf((*MockDBInterface)(nil).GetCategoryChildren), categoryId)
}

// GetChannelByDeliverectLinkId mocks base method.
func (m *MockDBInterface) GetChannelByDeliverectLinkId(deliverectLinkId string) (models.DeliverectChannel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChannelByDeliverectLinkId", deliverectLinkId)
	ret0, _ := ret[0].(models.DeliverectChannel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChannelByDeliverectLinkId indicates an expected call of GetChannelByDeliverectLinkId.
func (mr *MockDBInterfaceMockRecorder) GetChannelByDeliverectLinkId(deliverectLinkId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannelByDeliverectLinkId", reflect.TypeOf((*MockDBInterface)(nil).GetChannelByDeliverectLinkId), deliverectLinkId)
}

// GetChannelByStoreId mocks base method.
func (m *MockDBInterface) GetChannelByStoreId(storeId uuid.UUID) (models.DeliverectChannel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChannelByStoreId", storeId)
	ret0, _ := ret[0].(models.DeliverectChannel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChannelByStoreId indicates an expected call of GetChannelByStoreId.
func (mr *MockDBInterfaceMockRecorder) GetChannelByStoreId(storeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannelByStoreId", reflect.TypeOf((*MockDBInterface)(nil).GetChannelByStoreId), storeId)
}

// GetOpenStore mocks base method.
func (m *MockDBInterface) GetOpenStore(storeId uuid.UUID) (models.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOpenStore", storeId)
	ret0, _ := ret[0].(models.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOpenStore indicates an expected call of GetOpenStore.
func (mr *MockDBInterfaceMockRecorder) GetOpenStore(storeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpenStore", reflect.TypeOf((*MockDBInterface)(nil).GetOpenStore), storeId)
}

// GetOpenStores mocks base method.
func (m *MockDBInterface) GetOpenStores() ([]dtos.GetStoresOverviewResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOpenStores")
	ret0, _ := ret[0].([]dtos.GetStoresOverviewResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOpenStores indicates an expected call of GetOpenStores.
func (mr *MockDBInterfaceMockRecorder) GetOpenStores() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpenStores", reflect.TypeOf((*MockDBInterface)(nil).GetOpenStores))
}

// GetOpeningHours mocks base method.
func (m *MockDBInterface) GetOpeningHours(storeId uuid.UUID) ([]dtos.GetStoreOpeningHour, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOpeningHours", storeId)
	ret0, _ := ret[0].([]dtos.GetStoreOpeningHour)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOpeningHours indicates an expected call of GetOpeningHours.
func (mr *MockDBInterfaceMockRecorder) GetOpeningHours(storeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpeningHours", reflect.TypeOf((*MockDBInterface)(nil).GetOpeningHours), storeId)
}

// GetOrderById mocks base method.
func (m *MockDBInterface) GetOrderById(orderId uuid.UUID) (models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderById", orderId)
	ret0, _ := ret[0].(models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderById indicates an expected call of GetOrderById.
func (mr *MockDBInterfaceMockRecorder) GetOrderById(orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderById", reflect.TypeOf((*MockDBInterface)(nil).GetOrderById), orderId)
}

// GetOrderByStripeOrderId mocks base method.
func (m *MockDBInterface) GetOrderByStripeOrderId(stripeOrderId string) (models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByStripeOrderId", stripeOrderId)
	ret0, _ := ret[0].(models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByStripeOrderId indicates an expected call of GetOrderByStripeOrderId.
func (mr *MockDBInterfaceMockRecorder) GetOrderByStripeOrderId(stripeOrderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByStripeOrderId", reflect.TypeOf((*MockDBInterface)(nil).GetOrderByStripeOrderId), stripeOrderId)
}

// GetOrderItems mocks base method.
func (m *MockDBInterface) GetOrderItems(orderId uuid.UUID) ([]models.OrderItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderItems", orderId)
	ret0, _ := ret[0].([]models.OrderItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderItems indicates an expected call of GetOrderItems.
func (mr *MockDBInterfaceMockRecorder) GetOrderItems(orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderItems", reflect.TypeOf((*MockDBInterface)(nil).GetOrderItems), orderId)
}

// GetOrdersByStoreId mocks base method.
func (m *MockDBInterface) GetOrdersByStoreId(storeId uuid.UUID) ([]dtos.GetOrdersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrdersByStoreId", storeId)
	ret0, _ := ret[0].([]dtos.GetOrdersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrdersByStoreId indicates an expected call of GetOrdersByStoreId.
func (mr *MockDBInterfaceMockRecorder) GetOrdersByStoreId(storeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrdersByStoreId", reflect.TypeOf((*MockDBInterface)(nil).GetOrdersByStoreId), storeId)
}

// GetOrdersByUserId mocks base method.
func (m *MockDBInterface) GetOrdersByUserId(userId uuid.UUID) ([]dtos.GetOrdersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrdersByUserId", userId)
	ret0, _ := ret[0].([]dtos.GetOrdersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrdersByUserId indicates an expected call of GetOrdersByUserId.
func (mr *MockDBInterfaceMockRecorder) GetOrdersByUserId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrdersByUserId", reflect.TypeOf((*MockDBInterface)(nil).GetOrdersByUserId), userId)
}

// GetProductChildren mocks base method.
func (m *MockDBInterface) GetProductChildren(parentProductId uuid.UUID) ([]uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductChildren", parentProductId)
	ret0, _ := ret[0].([]uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductChildren indicates an expected call of GetProductChildren.
func (mr *MockDBInterfaceMockRecorder) GetProductChildren(parentProductId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductChildren", reflect.TypeOf((*MockDBInterface)(nil).GetProductChildren), parentProductId)
}

// GetProducts mocks base method.
func (m *MockDBInterface) GetProducts(storeId uuid.UUID) ([]dtos.GetStoreProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts", storeId)
	ret0, _ := ret[0].([]dtos.GetStoreProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockDBInterfaceMockRecorder) GetProducts(storeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockDBInterface)(nil).GetProducts), storeId)
}

// GetProductsByPlu mocks base method.
func (m *MockDBInterface) GetProductsByPlu(plu string, storeId uuid.UUID) ([]uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsByPlu", plu, storeId)
	ret0, _ := ret[0].([]uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsByPlu indicates an expected call of GetProductsByPlu.
func (mr *MockDBInterfaceMockRecorder) GetProductsByPlu(plu, storeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsByPlu", reflect.TypeOf((*MockDBInterface)(nil).GetProductsByPlu), plu, storeId)
}

// GetStoreById mocks base method.
func (m *MockDBInterface) GetStoreById(storeId uuid.UUID) (models.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStoreById", storeId)
	ret0, _ := ret[0].(models.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStoreById indicates an expected call of GetStoreById.
func (mr *MockDBInterfaceMockRecorder) GetStoreById(storeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoreById", reflect.TypeOf((*MockDBInterface)(nil).GetStoreById), storeId)
}

// GetStoreByMerchantUserId mocks base method.
func (m *MockDBInterface) GetStoreByMerchantUserId(merchantUserId string) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStoreByMerchantUserId", merchantUserId)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStoreByMerchantUserId indicates an expected call of GetStoreByMerchantUserId.
func (mr *MockDBInterfaceMockRecorder) GetStoreByMerchantUserId(merchantUserId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoreByMerchantUserId", reflect.TypeOf((*MockDBInterface)(nil).GetStoreByMerchantUserId), merchantUserId)
}

// GetTopProducts mocks base method.
func (m *MockDBInterface) GetTopProducts(storeId uuid.UUID) ([]dtos.GetStoreProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopProducts", storeId)
	ret0, _ := ret[0].([]dtos.GetStoreProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopProducts indicates an expected call of GetTopProducts.
func (mr *MockDBInterfaceMockRecorder) GetTopProducts(storeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopProducts", reflect.TypeOf((*MockDBInterface)(nil).GetTopProducts), storeId)
}

// GetUserByFirebaseId mocks base method.
func (m *MockDBInterface) GetUserByFirebaseId(firebaseId string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByFirebaseId", firebaseId)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByFirebaseId indicates an expected call of GetUserByFirebaseId.
func (mr *MockDBInterfaceMockRecorder) GetUserByFirebaseId(firebaseId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByFirebaseId", reflect.TypeOf((*MockDBInterface)(nil).GetUserByFirebaseId), firebaseId)
}

// GetUserByUserId mocks base method.
func (m *MockDBInterface) GetUserByUserId(userId uuid.UUID) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUserId", userId)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUserId indicates an expected call of GetUserByUserId.
func (mr *MockDBInterfaceMockRecorder) GetUserByUserId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUserId", reflect.TypeOf((*MockDBInterface)(nil).GetUserByUserId), userId)
}

// IsWhitelisted mocks base method.
func (m *MockDBInterface) IsWhitelisted(identifier string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsWhitelisted", identifier)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsWhitelisted indicates an expected call of IsWhitelisted.
func (mr *MockDBInterfaceMockRecorder) IsWhitelisted(identifier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsWhitelisted", reflect.TypeOf((*MockDBInterface)(nil).IsWhitelisted), identifier)
}

// MarkOrderAsPaid mocks base method.
func (m *MockDBInterface) MarkOrderAsPaid(orderId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkOrderAsPaid", orderId)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkOrderAsPaid indicates an expected call of MarkOrderAsPaid.
func (mr *MockDBInterfaceMockRecorder) MarkOrderAsPaid(orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkOrderAsPaid", reflect.TypeOf((*MockDBInterface)(nil).MarkOrderAsPaid), orderId)
}

// SetIsOpen mocks base method.
func (m *MockDBInterface) SetIsOpen(storeId uuid.UUID, isOpen bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetIsOpen", storeId, isOpen)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetIsOpen indicates an expected call of SetIsOpen.
func (mr *MockDBInterfaceMockRecorder) SetIsOpen(storeId, isOpen interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIsOpen", reflect.TypeOf((*MockDBInterface)(nil).SetIsOpen), storeId, isOpen)
}

// SetStoreImageUrl mocks base method.
func (m *MockDBInterface) SetStoreImageUrl(storeId uuid.UUID, imageUrl string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStoreImageUrl", storeId, imageUrl)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStoreImageUrl indicates an expected call of SetStoreImageUrl.
func (mr *MockDBInterfaceMockRecorder) SetStoreImageUrl(storeId, imageUrl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStoreImageUrl", reflect.TypeOf((*MockDBInterface)(nil).SetStoreImageUrl), storeId, imageUrl)
}

// UpdateChannelStatus mocks base method.
func (m *MockDBInterface) UpdateChannelStatus(status dtos.ChannelStatus, storeId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateChannelStatus", status, storeId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateChannelStatus indicates an expected call of UpdateChannelStatus.
func (mr *MockDBInterfaceMockRecorder) UpdateChannelStatus(status, storeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateChannelStatus", reflect.TypeOf((*MockDBInterface)(nil).UpdateChannelStatus), status, storeId)
}

// UpdateOrderStatus mocks base method.
func (m *MockDBInterface) UpdateOrderStatus(orderId uuid.UUID, orderStatus models.OrderStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderStatus", orderId, orderStatus)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrderStatus indicates an expected call of UpdateOrderStatus.
func (mr *MockDBInterfaceMockRecorder) UpdateOrderStatus(orderId, orderStatus interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderStatus", reflect.TypeOf((*MockDBInterface)(nil).UpdateOrderStatus), orderId, orderStatus)
}

// UpdateProductSortOrder mocks base method.
func (m *MockDBInterface) UpdateProductSortOrder(childProductId uuid.UUID, sortOrder int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductSortOrder", childProductId, sortOrder)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProductSortOrder indicates an expected call of UpdateProductSortOrder.
func (mr *MockDBInterfaceMockRecorder) UpdateProductSortOrder(childProductId, sortOrder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductSortOrder", reflect.TypeOf((*MockDBInterface)(nil).UpdateProductSortOrder), childProductId, sortOrder)
}

// UpdateProductsSnooze mocks base method.
func (m *MockDBInterface) UpdateProductsSnooze(productIds []uuid.UUID, snooze bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductsSnooze", productIds, snooze)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProductsSnooze indicates an expected call of UpdateProductsSnooze.
func (mr *MockDBInterfaceMockRecorder) UpdateProductsSnooze(productIds, snooze interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductsSnooze", reflect.TypeOf((*MockDBInterface)(nil).UpdateProductsSnooze), productIds, snooze)
}
