package service

import (
	"context"
	pb "payment-service/generated/payment_service"
	"payment-service/storage/postgres"
)

type PaymentService struct{
	pb.UnimplementedPaymentServiceServer
	Payment postgres.PaymentRepo
}

func NewPaymentService(payment postgres.PaymentRepo) *PaymentService{
	return &PaymentService{Payment: payment}
}

func (s *PaymentService) CreatePayment(ctx context.Context,payment *pb.CreatePaymentRequest)(*pb.CreatePaymentResponse,error){
	res,err := s.Payment.CreatePayment(payment)
	if err != nil{
		return nil,err
	}
	return res,nil
}

func (s *PaymentService) GetPayment(ctx context.Context,payment *pb.GetPaymentRequest)(*pb.GetPaymentResponse,error){
	res,err := s.Payment.GetPayment(payment)
	if err != nil{
		return nil,err
	}
	return res,nil
}

func (s *PaymentService) UpdatePayment(ctx context.Context,updatePayment *pb.UpdatePaymentRequest)(*pb.UpdatePaymentResponse,error){
	res,err := s.Payment.UpdatePayment(updatePayment)
	if err != nil{
		return nil,err
	}
	return res,nil
}