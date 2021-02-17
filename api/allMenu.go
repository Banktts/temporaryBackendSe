package api

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"context"
)



func AllMenu() {
	ExtraMenu:=extraMenu(17)
	fmt.Println(ExtraMenu)
}

func extraMenu(mid int) []bson.M {
	fmt.Println("allMenu")
	extra := connectMongoDB("AgilestNoRelationDB").Collection("EXTRA_ITEM")
	var extras []bson.M
	rawData,err:=extra.Find(context.TODO(),bson.M{"M_id": mid})
	if err != nil {
		fmt.Println(err)
	}
	rawData.All(context.TODO(),&extras)
	disconectMongoDB()
	return extras
}
