package main

import (
	"context"
	"fmt"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type UserInfo struct {
	UserName   string
	Age        int
	EmployeeId int
}

func main() {

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	start := time.Now()
	for i := 0; i < 100000; i++ {
		collection.InsertOne(ctx, bson.M{"name": fmt.Sprintf("pi %d", i), "value": math.Pi})
	}
	end := time.Since(start)
	fmt.Println(fmt.Sprintf("Elapsed: %s", end)) // Elapsed: 9.635583918s

}
