package postgres

import (
	"database/sql"
	pb "payment-service/generated/payment_service"
)

type PaymentRepo struct {
	DB *sql.DB
}

func NewPaymentRepo(db *sql.DB) *PaymentRepo {
	return &PaymentRepo{DB: db}
}

func (p *PaymentRepo) CreatePayment(payment *pb.CreatePaymentRequest)(*pb.CreatePaymentResponse,error){
	paym := pb.CreatePaymentResponse{}
	err := p.DB.QueryRow(`
		INSERT INTO Payments(
			reservation_id,
			amount,
			payment_method,
			payment_status
		)
		VALUES(
			$1,
			$2,
			$3,
			$4
		)
		returning
			id,
			reservation_id,
			amount,
			payment_method,
			payment_status`,
		payment.ReservationId,payment.Amount,payment.PaymentMethod,payment.PaymentStatus).Scan(
			paym.Payment.Id,paym.Payment.ReservationId,paym.Payment.Amount,paym.Payment.PaymentMethod,paym.Payment.PaymentStatus,
		)		
	if err != nil{
		return nil,err
	}
	return &paym,err
}

func (p *PaymentRepo) GetPayment(payment *pb.GetPaymentRequest)(*pb.GetPaymentResponse,error){
	resPayment := pb.GetPaymentResponse{}
	err := p.DB.QueryRow(`
		SELECT
			id,
			reservation_id,
			amount,
			payment_method,
			payment_status
		FROM
			Payments
		WHERE
			id = $1
		`,
		payment.Id).Scan(
			resPayment.Payment.Id,resPayment.Payment.ReservationId,resPayment.Payment.Amount,resPayment.Payment.PaymentMethod,resPayment.Payment.PaymentStatus,
		)
	if err != nil{
		return nil,err
	}
	return &resPayment,nil
}

func (p *PaymentRepo) UpdatePayment(updatePayment *pb.UpdatePaymentRequest)(*pb.UpdatePaymentResponse,error){
	resPayment := pb.UpdatePaymentResponse{}
	err := p.DB.QueryRow(`UPDATE
		Payments
		SET
			reservation_id = $1,
			amount = $2,
			payment_method = $3,
			payment_status = $4
		WHERE
			id = $5
		returning
			id,
			reservation_id,
			amount,
			payment_method,
			payment_status
		`,
		updatePayment.ReservationId,updatePayment.Amount,updatePayment.PaymentMethod,updatePayment.PaymentStatus,updatePayment.Id,
	).Scan(
		resPayment.Payment.Id,resPayment.Payment.ReservationId,resPayment.Payment.Amount,resPayment.Payment.PaymentMethod,resPayment.Payment.PaymentStatus,
	)
	if err != nil{
		return nil,err
	}
	return &resPayment,err
}