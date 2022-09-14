package deliverect

// Payment Model Type
const (
	CREDIT_CARD_ONLINE = iota
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

// Order Type
const (
	PICKUP = iota + 1
	DELIVERY
	EAT_IN
	CURBSIDE
)
