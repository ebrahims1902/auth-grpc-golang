package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodb struct {
	Client *mongo.Client
	Db     *mongo.Database
	Coll   *mongo.Collection
	ctx    context.Context
}

func (mgo *Mongodb) ConnMongo() (*Mongodb, context.Context) {

	mgo.Client, _ = mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017/task-manager"))

	mgo.ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)

	defer mgo.Client.Connect(mgo.ctx)
	mgo.Db = mgo.Client.Database("quickstart")
	mgo.Coll = mgo.Db.Collection("podcasts")
	return mgo, mgo.ctx
}
func main() {
	strmgo := new(Mongodb)
	mgo, ctx := strmgo.ConnMongo()
	fmt.Println("!!!!!!!", mgo)

	podcastsResult, err := mgo.Coll.InsertOne(ctx, bson.D{
		{"title", "The polyglot Develloper Podcast"},
		{"author", "Nic Raboy"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(podcastsResult.InsertedID)
}
