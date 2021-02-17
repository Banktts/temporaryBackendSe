package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"strconv"
)

func main() {
	

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("AgilestNoRelationDB").Collection("EXTRA_ITEM")
	if err != nil {
		fmt.Println(err)
	}
	csvFile, err := os.Open("mongo/extra_item.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		EID, _ := strconv.Atoi(line[0])
		PRICE, _ := strconv.Atoi(line[2])
		MID, _ := strconv.Atoi(line[3])
		insertResult, err := collection.InsertOne(context.TODO(), bson.D{{"E_id",EID},{"M_id",MID},{"E_name",line[1]},{"E_price",PRICE},{"E_type",line[4]} })
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	}

	err = client.Disconnect(context.TODO())
	fmt.Println("Connection to MongoDB closed.")
}
