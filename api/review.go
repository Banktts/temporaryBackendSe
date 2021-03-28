package api

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

type Review struct {
	Driver Driver `json:"Driver"`
	Restaurant Restaurant `json:"Restaurant"`
}

type Driver struct{
	D_id int 
	Star float64 
	Comment string 
}
type Restaurant struct{
	R_id int
	Star float64 
	Comment string 
}

func AddReview ( review Review ){
	AddDriverReview(review.Driver)
	AddRestaurantReview(review.Restaurant)
}

func AddDriverReview(driver Driver){
	extra := connectMongoDB().Collection("Driver_Review")
	AddOnLine:=bson.D{
		{Key: "D_ID", Value:driver.D_id},
		{Key: "star",Value:driver.Star},
		{Key: "comment",Value:driver.Comment},

	}
	_,err:= extra.InsertOne(context.TODO(),AddOnLine)
	if err != nil {
		fmt.Println(err)
	}
	disconectMongoDB()
}

func AddRestaurantReview(restaurant Restaurant){
	extra := connectMongoDB().Collection("Restaurant_Review")
	AddOnLine:=bson.D{
		{Key: "R_ID", Value:restaurant.R_id},
		{Key: "star",Value:restaurant.Star},
		{Key: "comment",Value:restaurant.Comment},

	}
	_,err:= extra.InsertOne(context.TODO(),AddOnLine)
	if err != nil {
		fmt.Println(err)
	}
	disconectMongoDB()
}