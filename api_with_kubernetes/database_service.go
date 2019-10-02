package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	HOST           = "localhost"
	PORT           = 27017
	DBNAME         = "service_db"
	COLLECTIONNAME = "user_info"
)

type UserInfo struct {
	UserName   string
	Age        int
	EmployeeId int
}

type DatabaseService interface {
	Save(*UserInfo)
	Retrieve(int) (*UserInfo, error)
}

type DatabaseServiceImpl struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewDatabaseServiceImp(host string, port int, dbname string, collection_name string) *DatabaseServiceImpl {
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port)))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Error : " + err.Error())
	}

	collection := client.Database(dbname).Collection(collection_name)
	return &DatabaseServiceImpl{client: client, collection: collection}
}

func (m *DatabaseServiceImpl) Save(user_info *UserInfo) error {
	insertResult, err := m.collection.InsertOne(context.TODO(), user_info)
	if err != nil {
		return err
	}
	fmt.Println("Inserted document : ", insertResult.InsertedID)
	return nil
}

func (m *DatabaseServiceImpl) Retrieve(id int) (*UserInfo, error) {
	var userInfo UserInfo
	err := m.collection.FindOne(context.TODO(), bson.M{"employeeid": id}).Decode(&userInfo)
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}
