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
	C_longitude  float64
	R_id         int
	Created_at   time.Time
	Ordertline   []Ordertline
}

type Ordertline struct {
	M_id        int      
	M_amount    int   
	M_comment	string      
	M_Extra 	[]M_Extra 
}
type Orderline struct {
	M_id        int      
	M_amount    int
	M_comment	string      
	M_Extra 	[]bson.M
}

type M_Extra struct{
	M_id int
	E_id int
}
type CustomerStruct struct{
	C_id int `Bson:"C_id"`
	C_name string `Bson:"C_name"`
	C_Tel string `Bson:"C_Tel"`
	C_latitude float64 `Bson:"C_latitude"`
	C_longtitude float64 `Bson:"C_longtitude"`
	R_id int `Bson:"R_id"`
	D_id int `Bson:"D_id"`
	O_id int `Bson:"O_id"`
	Created_at  time.Time  `Bson:"D_id"`
	Orderline   []Orderline `Bson:"Orderline "`
	DeliveryFee int `Bson:"DeliveryFee"`
	TotalPrice int `Bson:"TotalPrice"`
}
func AddOrder (order Order )  CustomerStruct {
	
	db := connectSqlDB()

	rows, err:= db.Query("select R_latitude ,R_longitude from restaurant where R_ID = ?",order.R_id)
	if err!=nil{
		fmt.Println(err)
	}
	var R_latitude float64
	var R_longitude float64
	for rows.Next() {
		

		err := rows.Scan(&R_latitude,&R_longitude)
		if err != nil {
			panic(err.Error())
		}
	}

	rows1, err1 := db.Query(" select (sqrt(power((D_latitude-?),2)+power((D_longitude-?),2)))*100  as distance  ,D_ID from delivery_man ORDER BY distance DESC limit 1"   ,R_latitude,R_longitude)
	if err1!=nil{
		fmt.Println(err1)
	}
	var distance float64
	var D_id int

	for rows1.Next() {
		
		err := rows1.Scan(&distance,&D_id)
		if err != nil {
			panic(err.Error())
		}
	}
	
	stmt, es := connectSqlDB().Prepare("INSERT INTO ordert (C_ID,R_ID,D_ID,created_at) VALUES (?,?,?,?) ")		
	if es != nil {
		panic(es.Error())
	}
	_,err2 := stmt.Exec(order.C_id,order.R_id,D_id,order.Created_at)
	if err2 != nil {
		panic(err2.Error())
	}

	var order_id int
	var totalPrice int 
	totalPrice = 0 
	err5 := db.QueryRow("select O_ID from ordert ORDER BY O_ID DESC LIMIT 1").Scan(&order_id)
	if err5 != nil {
	panic(err5.Error())
	}
	fmt.Println("order",order_id)

	
	rows4, err4:= db.Query(" select menu.M_price,orderline.amount from orderline natural join menu where O_ID = ?",order_id-1)
	if err4!=nil{
		fmt.Println(err4)
	}
	
	for rows4.Next() {
		var M_price int
		var amount int
		err4 := rows4.Scan(&M_price, &amount)
		if err4 != nil {
			panic(err4.Error())
		}
		fmt.Println(M_price , amount )
		totalPrice  = totalPrice + M_price * amount 
	}
	

	totalPrice = totalPrice+10

	customer  := CustomerStruct{
		C_id   : order.C_id,
		C_name : order.C_name ,
		C_Tel : order.C_tel,
		C_latitude   : order.C_latitude,
		C_longtitude : order.C_longitude,
		R_id : order.R_id,
		D_id : D_id,
		O_id : order_id,
		Created_at : order.Created_at,
		Orderline   : AddOrderline(order.Ordertline),
		DeliveryFee : 10,
		TotalPrice : totalPrice,
	}
	
	

	return customer


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

		rows, err := db.Query(`INSERT INTO orderline (O_ID, M_ID, amount, special_inst) VALUES (?, ?, ?, ?)`, order_id, s.M_id, s.M_amount, s.M_comment)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for _,m := range s.M_Extra{
			MExtraAdd(order_id, m.E_id, m.M_id)
		}
		odl := Orderline{
			M_id:        s.M_id,
			M_amount:    s.M_amount,
			M_comment:	s.M_comment,
			M_Extra: MExtra(order_id),
		}		
		allOrderline = append(allOrderline,odl)		
		
		
		
	}

	return allOrderline

}