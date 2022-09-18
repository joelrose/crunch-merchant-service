package deliverect

type PaymentType int

const (
	CREDIT_CARD_ONLINE PaymentType = iota
	CASH
	ON_DELIVERY
	ONLINE
	CREDIT_CARD_AT_DOOR
	PIN_AT_DOOR
	VOUCHER_AT_DOO
	CHEQUE
	BANK_CONTACT
	OTHER
)

type OrderType int

const (
	PICKUP OrderType = iota + 1
	DELIVERY
	EAT_IN
	CURBSIDE
)

type AccountType int

const (
	PARTNER AccountType = iota + 1
	CHAIN
	CUSTOMER
)
