package payment

import (
    "context"

    "google.golang.org/grpc"
)

// Minimal stub types to satisfy imports until real generated code is available.
type CreatePaymentRequest struct {
    UserId     int64
    OrderId    int64
    TotalPrice float32
}

type CreatePaymentResponse struct {
    PaymentId int64
    BillId    int64
}

type PaymentClient interface {
    Create(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error)
}

type paymentClient struct{
    conn *grpc.ClientConn
}

func NewPaymentClient(conn *grpc.ClientConn) PaymentClient {
    return &paymentClient{conn: conn}
}

func (c *paymentClient) Create(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error) {
    // stub: in real generated code this will perform the RPC.
    return &CreatePaymentResponse{PaymentId: 1, BillId: 1}, nil
}
