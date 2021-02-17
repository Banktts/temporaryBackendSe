package api

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)



func AllMenu()[]bson.M {
	ExtraMenu:=extraMenu(17)
	return ExtraMenu

}
func extraMenu(mid int) []bson.M {
	fmt.Println("allMenu")
	extra := connectMongoDB().Collection("EXTRA_ITEM")
	var extras []bson.M
	rawData,err:=extra.Find(context.TODO(),bson.M{"M_id": mid})
	if err != nil {
		fmt.Println(err)
	}
	rawData.All(context.TODO(),&extras)
	disconectMongoDB()
	return extras
}
