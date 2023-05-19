package Model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	StudentId    string             `json:"studentid" bson:"studentid,omitempty"`
	FirstName    string             `json:"firstname" bson:"firstname,omitempty"`
	LastName     string             `json:"lastname" bson:"lastname,omitempty"`
	University   string             `json:"universityid" bson:"universityid,omitempty"`
	Password     string             `json:"password" bson:"password,omitempty"`
	CurrentMoney int                `json:"currentmoney" bson:"currentmoney,omitempty"`
}
type Report struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Reporter       string             `json:"reporter" bson:"reporter,omitempty"`
	Reportee       string             `json:"reportee" bson:"reportee,omitempty"`
	ReportedCoupon Coupon             `json:"reportedcoupon" bson:"reportedcoupon,omitempty"`
}
type Coupon struct {
	ID         int64  `json:"_id,omitempty" bson:"_id,omitempty"`
	Price      int    `json:"price" bson:"price,omitempty"`
	StudentId  string `json:"studentid" bson:"studentid,omitempty"`
	University string `json:"universityid" bson:"universityid,omitempty"`
	CanteenID  string `json:"canteenid" bson:"canteenid,omitempty"`
	Code       string `json:"code" bson:"code,omitempty"`
	FoodName   string `json:"foodname" bson:"foodname,omitempty"`
}
type University struct {
	ID string `json:"universityid" bson:"universityid"`
}
type Canteen struct {
	ID           string `json:"canteenid" bson:"canteenid"`
	UniversityID string `json:"universityid" bson:"universityid"`
}
type Admin struct {
	ID       string `json:"adminid" bson:"adminid"`
	UserName string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Name     string `json:"name" bson:"name"`
}
type CurStatus struct {
	Stat string `json:"stat" bson:"stat"`
}
type DeletingCoupon struct {
	ID                int64  `json:"_id,omitempty" bson:"_id,omitempty"`
	BuyerStudentID    string `json:"buyersid" bson:"buyersid"`
	BuyerUniversityID string `json:"buyeruid" bson:"buyeruid"`
}
