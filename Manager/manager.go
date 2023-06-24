package Manager

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var CounterCollection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	CounterCollection = client.Database("food").Collection("counter")
}

func GetNewCouponId() int64 {
	id := "counter"
	filter := bson.M{"_id": id}
	var result Cnt
	err := CounterCollection.FindOne(context.Background(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		var tmp Cnt
		tmp.Cnt = 0
		tmp.ID = "counter"
		_, err := CounterCollection.InsertOne(context.Background(), tmp)
		if err != nil {
			log.Fatal(err)
		}
		return tmp.Cnt
	}
	filter2 := bson.M{"_id": result.ID}
	update2 := bson.M{"$set": bson.M{"cnt": result.Cnt + 1}}
	res2, err := CounterCollection.UpdateOne(context.Background(), filter2, update2)
	if err != nil {
		log.Fatal(res2)
		log.Fatal(err)
	}
	return result.Cnt + 1
}

type Cnt struct {
	ID  string `json:"_id,omitempty" bson:"_id,omitempty"`
	Cnt int64  `json:"cnt,omitempty" bson:"cnt,omitempty"`
}
