package paymentRepository

import (
	"gorm.io/gorm"
	"tinderleaked/modules/payment"
)

type sqlRepository struct {
	Conn *gorm.DB
}

func NewPaymentRepository(Conn *gorm.DB) payment.RepositoryPayment {
	return &sqlRepository{Conn: Conn}
}
