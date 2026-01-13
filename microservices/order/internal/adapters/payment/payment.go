package payment

import (
	"context"

	pb "github.com/Mellanie-Marques/microservices-proto/golang/payment/payment"
	"github.com/Mellanie-Marques/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	client pb.PaymentClient // from the generated protobuf code
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	client := pb.NewPaymentClient(conn)
	return &Adapter{client: client}, nil
}

func (a *Adapter) Charge(order *domain.Order) error {
	_, err := a.client.Create(context.Background(), &pb.CreatePaymentRequest{
		UserId:     order.CustomerID,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	})
	return err
}
