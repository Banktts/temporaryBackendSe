package api

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	_ "github.com/go-sql-driver/mysql"
)
type Order struct {
	C_id         int
	C_name       string
	C_tel        string
	C_latitude   float64
	C_longtitude float64
	R_id         int
	Created_at   time.Time
	Ordertline   []Ordertline
}

type Ordertline struct {
	M_id        int      
	M_amount    int      
	M_Extra 	[]M_Extra 
}
type Orderline struct {
	M_id        int      
	M_amount    int      
	M_Extra 	[]bson.M
}

type M_Extra struct{
	M_id int
	E_id int
}


func MExtraAdd(O_id int,E_id int,M_id int)   {
	fmt.Println("MExtraAdd")
	extra := connectMongoDB().Collection("EXTRA_Add_On")
	AddOnLine:=bson.D{
		{Key: "O_id", Value:O_id},
		{Key: "E_id",Value:E_id},
		{Key: "M_id",Value:M_id},

	}
	_,err:= extra.InsertOne(ctx,AddOnLine)
	if err != nil {
		fmt.Println(err)
	}
	disconectMongoDB()

}

func MExtra(O_id int) []bson.M{
	fmt.Println("MExtra")
	extra := connectMongoDB().Collection("EXTRA_Add_On")
	var extras []bson.M
	rawData, err := extra.Find(context.TODO(), bson.M{"O_id": O_id})
	if err != nil {
		fmt.Println(err)
	}
	rawData.All(context.TODO(), &extras)
	disconectMongoDB()
	return extras
}

func AddOrderline(o_list []Ordertline) []Orderline {

	db := connectSqlDB()

	var order_id int
	err := db.QueryRow("select O_ID from ordert ORDER BY O_ID DESC LIMIT 1").Scan(&order_id)
	if err != nil {
		panic(err.Error())
	}
	var allOrderline []Orderline

	for _, s := range o_list  {

		rows, err := db.Query(`INSERT INTO orderline (O_ID, M_ID, amount) VALUES (?, ?, ?)`, order_id, s.M_id, s.M_amount)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		odl := Orderline{
			M_id:        s.M_id,
			M_amount:    s.M_amount,
			M_Extra: MExtra(order_id),
		}		
		allOrderline = append(allOrderline,odl)		
		
		for _,m := range s.M_Extra{
			MExtraAdd(order_id, m.E_id, m.M_id)
		}
		
	}

	return allOrderline

}