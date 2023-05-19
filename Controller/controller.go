package Controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

)

const connectionString = "mongodb+srv://avatar:{password}@cluster0.ibwwj5y.mongodb.net/test"
const dbName = "dblab"
const colName = "id"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("collection is ready")
}

func insertOneQuery() {
}

func updateWebSite() {
}
func deleteOneWebSite() {
}
func deleteAll() int64 {
}

func getAllWebsites() []primitive.M {
}

func GetAllWebsitesJSON(w http.ResponseWriter, r *http.Request) {
}

func CreateWebsite(w http.ResponseWriter, r *http.Request) {
}

func MarkAsChecked(w http.ResponseWriter, r *http.Request) {
}

func DeleteOneWebsite(w http.ResponseWriter, r *http.Request) {
}
func DeleteAllWebsite(w http.ResponseWriter, r *http.Request) {
}
