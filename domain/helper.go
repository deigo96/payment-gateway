package domain

import (
	"log"
	"strconv"
	"time"

	"github.com/midtrans/midtrans-go/coreapi"
)

func RespnsePayment (c *coreapi.ChargeResponse, bank string)*PreTranscations{
	var response PreTranscations

	switch c.PaymentType {
	case "bank_transfer":
		if bank == "BCA" || bank == "BNI" {
			response.PaymentCode = &c.VaNumbers[0].VANumber
		} else if bank == "PER" || bank == "ATM" {
			response.PaymentCode = &c.PermataVaNumber
		}
		return &response
	case "echannel":
		response.PaymentCode = &c.BillerCode
		return &response
	case "gopay", "qris":
		response.PaymentCode = &c.Actions[0].URL
		return &response
	case "cstore":
		response.PaymentCode = &c.PaymentCode
		return &response
	default:
		return nil
	}

}


func StringToTime(str string) time.Time{
	date, error := time.Parse("2006-01-02 15:04:05", str)
  
	if error != nil {
		log.Println(error)
		return date
	}

	return date
}


func StringToInt(str string) int {
	i, _ := strconv.Atoi(str)

	return i
}

func IntToString(i int) string {
	str := strconv.Itoa(i)

	return str
}

func TimeToString(t time.Time) string {
	str := t.Format("2006-01-02 15:04:05")
	return str
}

func OrderId() string {
	t := time.Now()
	order := t.UnixMilli()
	orderId := "INV-" + strconv.FormatInt(order, 10)

	return orderId
}
