package api

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type Menu struct {
	R_id        int      `Bson:"R_id"`
	M_id        int      `Bson:"M_id"`
	M_name      string   `Bson:"M_name"`
	M_price     int      `Bson:"M_price"`
	M_image_url string   `Bson:"M_image_url"`
	M_Extra     []bson.M `Bson:"M_Extra"`
}

func AllMenu(R_id int) []Menu {
	var allmenu []Menu
	fmt.Println("serrrrrrrr")
	rows, err := connectSqlDB().Query("select * from menu")
	fmt.Println(rows)
	if err != nil {
		panic(err.Error())
	}
	//defer rows.Close()

	for rows.Next() {
		fmt.Println("loop")
		var R_id int
		var M_id int
		var M_name string
		var M_price int
		var M_image_url string

		err := rows.Scan(&R_id, &M_id, &M_name, &M_price, &M_image_url)

		if err != nil {
			panic(err.Error())
		}
		menu := Menu{
			R_id:        R_id,
			M_id:        M_id,
			M_name:      M_name,
			M_price:     M_price,
			M_image_url: M_image_url,

			M_Extra: extraMenu(17),
		}
		allmenu = append(allmenu, menu)

	}
	fmt.Println("setftesxdrgd")
	fmt.Println(allmenu)
	return allmenu

}
func extraMenu(mid int) []bson.M {
	fmt.Println("allMenu")
	extra := connectMongoDB().Collection("EXTRA_ITEM")
	var extras []bson.M
	rawData, err := extra.Find(context.TODO(), bson.M{"M_id": mid})
	if err != nil {
		fmt.Println(err)
	}
	rawData.All(context.TODO(), &extras)
	disconectMongoDB()
	return extras
}
