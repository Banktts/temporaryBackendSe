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

type M_Extra struct{
	M_id int
	E_id int
	Comment string
}


type Orderline struct {
	O_id        int      `Bson:"O_id"`
	M_id        int      `Bson:"M_id"`
	amount        int      `Bson:"amount"`
	special_inst string `Bson:"special_inst"`
}

func MExtraAdd(O_id int,E_id int,M_id int,Comment string) []bson.M  {
	fmt.Println("MExtraAdd")
	extra := connectMongoDB().Collection("EXTRA_Add_On")
	AddOnLine:=bson.D{
		{Key: "O_id", Value:O_id},
		{Key: "E_id",Value:E_id},
		{Key: "M_id",Value:M_id},
		{Key: "Comment",Value:Comment},
	}
	_,err:= extra.InsertOne(ctx,AddOnLine)
	if err != nil {
		fmt.Println(err)
	}
	disconectMongoDB()
	return MExtra(O_id)
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

// func AddOrderline() {
// 	db := connectSqlDB()
// 	sqlStatement := `INSERT INTO orderline (O_ID, M_ID, E_ID, special_inst) VALUES ($1, $2, $3, $4)`
// 	res, err := db.Exec(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun").Scan(&id)
// 	if err != nil {
// 	  panic(err)
// 	}
// }