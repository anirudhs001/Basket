package models

import "errors"

//ErrSellerAlreadyExists constant
var ErrSellerAlreadyExists error

func init() {

	ErrSellerAlreadyExists = errors.New("err: seller already exists")
}
