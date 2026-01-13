package ports

import "github.com/Mellanie-Marques/microservices/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(order *domain.Order) error
}
