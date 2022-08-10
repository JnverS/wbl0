package model

import (
	"errors"
	"reflect"

	"github.com/hashicorp/go-multierror"
)

type Order struct {
	OrderUid    string `json:"order_uid"`
	TrackNumber string `json:"track_number"`
	Entry       string `json:"entry"`
	Delivery    struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		Address string `json:"address"`
		Region  string `json:"region"`
		Email   string `json:"email"`
	} `json:"delivery"`
	Payment struct {
		Transaction  string `json:"transaction"`
		RequestId    string `json:"request_id"`
		Currency     string `json:"currency"`
		Provider     string `json:"provider"`
		Amount       int    `json:"amount"`
		PaymentDt    int    `json:"payment_dt"`
		Bank         string `json:"bank"`
		DeliveryCost int    `json:"delivery_cost"`
		GoodsTotal   int    `json:"goods_total"`
		CustomFee    int    `json:"custom_fee"`
	} `json:"payment"`
	Items []struct {
		ChrtId      int    `json:"chrt_id"`
		TrackNumber string `json:"track_number"`
		Price       int    `json:"price"`
		Rid         string `json:"rid"`
		Name        string `json:"name"`
		Sale        int    `json:"sale"`
		Size        string `json:"size"`
		TotalPrice  int    `json:"total_price"`
		NmId        int    `json:"nm_id"`
		Brand       string `json:"brand"`
		Status      int    `json:"status"`
	} `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmId              int       `json:"sm_id"`
	DateCreated       string `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

func (o Order) Validate() error {
	var err error
	if o.OrderUid == "" {
		err = multierror.Append(err, errors.New("empty Order UID"))
	}
	if o.TrackNumber == "" {
		err = multierror.Append(err, errors.New("incorrect Tracknumber"))
	}
	if o.Entry == "" {
		err = multierror.Append(err, errors.New("incorrect Entry"))
	}
	if o.Delivery.Name == "" {
		err = multierror.Append(err, errors.New("incorrect Delivery name"))
	}
	if o.Payment.Transaction == "" {
		err = multierror.Append(err, errors.New("incorrect Payment Transaction"))
	}
	if reflect.ValueOf(o.Items).IsNil() {
		err = multierror.Append(err, errors.New("incorrect Items"))
	}
	if o.Locale == "" {
		err = multierror.Append(err, errors.New("incorrect Locale"))
	}
	if o.CustomerId == "" {
		err = multierror.Append(err, errors.New("incorrect CustomerId"))
	}
	if o.DeliveryService == "" {
		err = multierror.Append(err, errors.New("incorrect DeliveryService"))
	}
	if o.Shardkey == "" {
		err = multierror.Append(err, errors.New("incorrect Shardkey"))
	}
	if o.DateCreated == "" {
		err = multierror.Append(err, errors.New("incorrect DateCreated"))
	}
	if o.OofShard == "" {
		err = multierror.Append(err, errors.New("incorrect OofShard"))
	}

	return err
}
