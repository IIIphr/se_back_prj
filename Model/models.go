package Model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	StudentId    string             `json:"studentid"`
	Password     string             `json:"password"`
	History      UserHistory        `json:"userhistory"`
	CurrentMoney int                `json:"currentmoney"`
}
type UserHistory struct {
	History []Coupon `json:"history"`
}
type Report struct {
	Report         string `json:"report"`
	Reporter       string `json:"reporter"`
	Reportee       string `json:"reportee"`
	ReportedCoupon Coupon `json:"reportedcoupon"`
}
type Coupon struct {
	Number int  `json:"number"`
	Price  int  `json:"price"`
	Owner  User `json:"user"`
}
