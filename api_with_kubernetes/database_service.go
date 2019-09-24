package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	HOST           = "localhost"
	PORT           = 27017
	DBNAME         = "service_db"
	COLLECTIONNAME = "user_info"
)

type DatabaseService interface {
	Save(UserInfo)
	Retrieve(int) UserInfo
}

type DatabaseServiceImpl struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (d *DatabaseServiceImpl) NewDatabaseServiceImp(host string, port int, dbname string, collection_name string) {
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port)))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Error : " + err.Error())
	}
	d.collection = client.Database(dbname).Collection(collection_name)
}

func (m DatabaseServiceImpl) Save(user_info UserInfo) {

}
