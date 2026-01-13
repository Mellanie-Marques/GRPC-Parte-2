package ports

import "github.com/Mellanie-Marques/microservices/order/internal/application/core/domain"

type APIPort interface {
    PlaceOrder(order domain.Order) (domain.Order, error)
}
