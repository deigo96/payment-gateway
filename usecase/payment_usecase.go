package usecase

import (
	"account-service/bootstrap"
	"account-service/domain"
	"encoding/json"
	"errors"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type paymentUsecase struct {
	paymentRepository domain.PaymentRepository
	contextTimeout time.Duration
	env bootstrap.Env
}

func NewPaymentUsecase(paymentRepository domain.PaymentRepository, timeout time.Duration, env bootstrap.Env) domain.PaymentUsecase {
	return &paymentUsecase{
		paymentRepository: paymentRepository,
		contextTimeout: timeout,
		env: env,
	}
}


func (p *paymentUsecase) StoreTransaction(data domain.StoreTranscation) (interface{}, error) {

	var chargeReq *coreapi.ChargeReq
	item := midtrans.ItemDetails{}
	items := []midtrans.ItemDetails{}
	orderId := domain.OrderId()
	total := 0

	for _, val := range data.Items {
		item.ID = val.Id
		item.Name = val.ItemName
		item.Qty = val.Quantity
		item.Price = val.Price
		total += int((val.Price * int64(val.Quantity)))

		items = append(items, item)

	}
	Customer := midtrans.CustomerDetails{
		Email: data.Customers.Email,
		FName: data.Customers.Username,
		Phone: data.Customers.Phone,
	}
	expiry := coreapi.CustomExpiry{
		OrderTime: time.Now().Format("2006-01-02 15:04:05 +0700"),
		ExpiryDuration: p.env.Expiry,
		Unit: "hour",
	}
	detail := midtrans.TransactionDetails{
		OrderID: orderId,
		GrossAmt: int64(total),
	}

	switch data.CodeBank {
	case "BCA":
		chargeReq = &coreapi.ChargeReq{
			PaymentType: coreapi.CoreapiPaymentType("bank_transfer"),
			TransactionDetails: detail,
			BankTransfer: &coreapi.BankTransferDetails{
				Bank:     midtrans.Bank("bca"),
			},
			Items: &items,
			CustomerDetails: &Customer,
			CustomExpiry: &expiry,
		}
	case "BNI":
		chargeReq = &coreapi.ChargeReq{
			PaymentType: coreapi.CoreapiPaymentType("bank_transfer"),
			TransactionDetails: detail,
			BankTransfer: &coreapi.BankTransferDetails{
				Bank:     midtrans.Bank("bni"),
			},
			Items: &items,
			CustomerDetails: &Customer,
			CustomExpiry: &expiry,
		}
	case "PER", "ATM":
		chargeReq = &coreapi.ChargeReq{
			PaymentType: coreapi.CoreapiPaymentType("bank_transfer"),
			TransactionDetails: detail,
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.Bank("permata"),
				Permata: &coreapi.PermataBankTransferDetail{
					RecipientName: data.Customers.Username,
				},
			},
			Items: &items,
			CustomerDetails: &Customer,
			CustomExpiry: &expiry,
		}
	case "MBP":
		chargeReq = &coreapi.ChargeReq{
			PaymentType: coreapi.CoreapiPaymentType("echannel"),
			TransactionDetails: detail,
			EChannel: &coreapi.EChannelDetail{
				BillInfo1: "Payment:",
				BillInfo2: "Online purchase",
				BillKey:   data.Customers.Phone,
			},
			Items: &items,
			CustomerDetails: &Customer,
			CustomExpiry: &expiry,
		}
	case "ALFA":
		chargeReq = &coreapi.ChargeReq{
			PaymentType: coreapi.CoreapiPaymentType("cstore"),
			TransactionDetails: detail,
			Items: &items,
			CustomerDetails: &Customer,
			ConvStore: &coreapi.ConvStoreDetails{
				Store:   "alfamart",
				Message: "Payment Gateway",
			},
			CustomExpiry: &expiry,
		}
	case "INDO":
		chargeReq = &coreapi.ChargeReq{
			PaymentType: coreapi.CoreapiPaymentType("cstore"),
			TransactionDetails: detail,
			Items: &items,
			CustomerDetails: &Customer,
			ConvStore: &coreapi.ConvStoreDetails{
				Store:   "indomaret",
				Message: "Payment Gateway",
			},
			CustomExpiry: &expiry,
		}
	case "QRIS":
		chargeReq = &coreapi.ChargeReq{
			PaymentType: coreapi.CoreapiPaymentType("qris"),
			TransactionDetails: detail,
			Items: &items,
			CustomerDetails: &Customer,
			Qris: &coreapi.QrisDetails{
				Acquirer: "gopay",
			},
			CustomExpiry: &expiry,
		}
	case "GPY":
		chargeReq = &coreapi.ChargeReq{
			PaymentType: coreapi.CoreapiPaymentType("gopay"),
			TransactionDetails: detail,
			Items: &items,
			CustomerDetails: &Customer,
			CustomExpiry: &expiry,
		}
	default:
		return nil,errors.New("invalid code bank")
	}

	reqCharge, err := json.Marshal(chargeReq)
	if err != nil {
		return nil, errors.New(domain.InternalServerError)
	}

	reqItems, err := json.Marshal(data.Items)
	if err != nil {
		return nil, errors.New(domain.InternalServerError)
	}

	bank, err := p.paymentRepository.GetBankByCode(data.CodeBank)
	if err != nil {
		return nil, err
	}

	lenItems := len(data.Items)
	gross := int64(total)

	dataTransaction := domain.PreTranscations{
		UserId: data.Customers.UserId,
		Items: reqItems,
		Status: 0,
		GrossAmount: &gross,
		RequestJson: reqCharge,
		CreatedAt: time.Now(),
		Quantity: int32(lenItems),
		OrderId: orderId,
		PaymentVia: &bank.Bank,
	}

	
	if err := p.paymentRepository.StoreTranscation(*chargeReq, dataTransaction); err != nil {
		return nil, err
	}


	return nil, nil
}

func (p *paymentUsecase) CancelTransaction(req domain.CancelParam) (*domain.CancelTransaction, error) {
	var record domain.CancelTransaction

	data, err := p.paymentRepository.CancelTransaction(req)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New(domain.RecordNotFound)
	}

	record.OrderId = data.OrderId
	record.PaymentVia = *data.PaymentVia
	record.StatusTransaction = *data.StatusTransaction

	return &record, nil
}

