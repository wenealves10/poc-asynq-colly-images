package model

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

type ProductStatus string

const (
	ProductStatusActive       ProductStatus = "ACTIVE"
	ProductStatusOutOfStock   ProductStatus = "OUT_OF_STOCK"
	ProductStatusPaused       ProductStatus = "PAUSED"
	ProductStatusUnavailable  ProductStatus = "UNAVAILABLE"
	ProductStatusRemoved      ProductStatus = "REMOVED"
	ProductStatusScrapeError  ProductStatus = "SCRAPE_ERROR"
)

type PaymentMethod string

const (
	PaymentMethodNone        PaymentMethod = "NONE"
	PaymentMethodPix         PaymentMethod = "PIX"
	PaymentMethodBoleto      PaymentMethod = "BOLETO"
	PaymentMethodCreditCard  PaymentMethod = "CREDIT_CARD"
)

type SubscriptionStatus string

const (
	SubscriptionStatusTrial    SubscriptionStatus = "TRIAL"
	SubscriptionStatusActive   SubscriptionStatus = "ACTIVE"
	SubscriptionStatusCanceled SubscriptionStatus = "CANCELED"
	SubscriptionStatusExpired  SubscriptionStatus = "EXPIRED"
)

type InvoiceStatus string

const (
	InvoiceStatusPending InvoiceStatus = "PENDING"
	InvoiceStatusPaid    InvoiceStatus = "PAID"
	InvoiceStatusFailed  InvoiceStatus = "FAILED"
)