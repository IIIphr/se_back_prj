package Model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	StudentId    string             `json:"studentid" bson:"studentid,omitempty"`
	Password     string             `json:"password" bson:"password,omitempty"`
	History      UserHistory        `json:"userhistory" bson:"userhistory,omitempty"`
	CurrentMoney int                `json:"currentmoney" bson:"currentmoney,omitempty"`
}
type UserHistory struct {
	History []Coupon `json:"history" bson:"history,omitempty"`
}
type Report struct {
	Report         string `json:"report" bson:"report,omitempty"`
	Reporter       string `json:"reporter" bson:"reporter,omitempty"`
	Reportee       string `json:"reportee" bson:"reportee,omitempty"`
	ReportedCoupon Coupon `json:"reportedcoupon" bson:"reportedcoupon,omitempty"`
}
type Coupon struct {
	Number int    `json:"number" bson:"number,omitempty"`
	Price  int    `json:"price" bson:"price,omitempty"`
	Owner  User   `json:"user" bson:"user,omitempty"`
	Self   string `json:"self" bson:"self,omitempty"`
}
