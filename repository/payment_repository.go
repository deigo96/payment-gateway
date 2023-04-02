package repository

import (
	"account-service/bootstrap"
	"account-service/domain"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
)

type paymentRepository struct {
	db *gorm.DB
	env bootstrap.Env
}

func NewPaymentRepository(db *gorm.DB, env bootstrap.Env) domain.PaymentRepository {
	return &paymentRepository{
		db: db,
		env: env,
	}
}

func (pr *paymentRepository) midtransClient() *coreapi.Client{
	var client = coreapi.Client{}
	key := pr.env.ServerKey
	envMidtans := pr.env.MidtransEnv
	client.New(key, midtrans.EnvironmentType(envMidtans))
	
	return &client
}

func (pr *paymentRepository) GetListBank() (record []domain.ListBank, err error) {
	if err := pr.db.Table("list_bank").Find(&record).Error; err != nil {
		return nil, errors.New(domain.InternalServerError)
	}

	return record, nil
}

func (pr *paymentRepository) GetBankByCode(code string) (record *domain.ListBank, err error) {
	if err := pr.db.Table("list_bank").Where("code_bank = ?", code).Scan(&record).Error; err != nil {
		return nil, errors.New(domain.InternalServerError)
	}

	return record, nil
}

func (pr *paymentRepository) StoreTranscation(coreapi coreapi.ChargeReq, data domain.PreTranscations) error {
	tx := pr.db.Begin()
	defer func ()  {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		log.Println("error store transcation: ", err)
		return errors.New(domain.InternalServerError)
	}

	err := tx.Table("pre_transaction").Create(&data)
	if err.Error != nil {
		log.Println("error store transcation: ",err.Error)
		tx.Rollback()
		return errors.New(domain.InternalServerError)
	}

	charge, errCharge := pr.midtransClient().ChargeTransaction(&coreapi)
	if errCharge != nil {
		log.Println("error charge transaction: ", errCharge)
		tx.Rollback()
		return errCharge
	}

	dataResponse := domain.RespnsePayment(charge, *data.PaymentVia)

	expTrx := domain.StringToTime(charge.TransactionTime).Add(time.Duration(pr.env.Expiry) * time.Hour)
	data.ExpiryAt = &expTrx
	data.TransactionId = &charge.TransactionID
	data.StatusTransaction = &charge.TransactionStatus
	data.PaymentCode = dataResponse.PaymentCode
	data.PaymentType = &charge.PaymentType
	
	res := tx.Table("pre_transaction").Where("order_id = ?", data.OrderId).Updates(&data)
	if res.Error != nil {
		log.Println("error store transcation: ",res.Error)
		tx.Rollback()
		return errors.New(domain.InternalServerError)
	}


	return tx.Commit().Error
}

func (pr *paymentRepository) CancelTransaction(req domain.CancelParam) (record *domain.PreTranscations, err error) {
	tx := pr.db.Begin()
	defer func ()  {
		if er := recover(); er != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		log.Println("error cancel transcation: ", err)
		return nil, errors.New(domain.InternalServerError)
	} 

	query := "order_id = '"+req.OrderId+"'"
	if req.OrderId == "" {
		query = "transaction_id = '"+req.TransactionId+"'"

	}

	if err := tx.Table("pre_transaction").Where(query).Updates(map[string]interface{}{
		"status": 2,
		"status_transaction": "canceled",
		"updated_at": time.Now(),
	}).Error; err != nil {
		log.Println("error updating transaction: ", err)
		tx.Rollback()
		return nil, errors.New(domain.InternalServerError)
	}

	cancel, errCancel := pr.midtransClient().CancelTransaction(req.OrderId)
	if errCancel != nil {
		log.Println("error cancel midtrans: ", err)
		tx.Rollback()
		return nil, errors.New(domain.InternalServerError)
	}

	fmt.Println(cancel)

	if err := tx.Table("pre_transaction").Where(query).Scan(&record).Error ; err != nil {
		log.Println("error get data: ", err)
		tx.Rollback()
		return nil, errors.New(domain.InternalServerError)
	}

	return record, tx.Commit().Error
}

