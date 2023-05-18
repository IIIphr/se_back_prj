package Model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WebSite struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	URL         string             `json:"url,omitempty"`
	WebsiteData WebSiteData        `json:"websitedata,omitempty"`
	Checked     bool               `json:"checked,omitempty"`
}

type WebSiteData struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
