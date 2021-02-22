package api

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

<<<<<<< HEAD
=======



>>>>>>> main
type Menu struct {
	R_id        int      `Bson:"R_id"`
	M_id        int      `Bson:"M_id"`
	M_name      string   `Bson:"M_name"`
	M_price     int      `Bson:"M_price"`
	M_image_url string   `Bson:"M_image_url"`
	M_Extra     []bson.M `Bson:"M_Extra"`
}

type Typ struct {
	T_name  string  `Bson:"T_name"`
	T_type  []Menu    `Bson:"T_type"`
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
<<<<<<< HEAD
=======
	
>>>>>>> main
	return extras
}

func addData(alltyp []Typ,typeFood string,R_id int) []Typ {
<<<<<<< HEAD
=======
	
>>>>>>> main
	stmt, err := connectSqlDB().Prepare("select * from menu where R_id = ? and M_type = ?")		
	rows,err := stmt.Query(R_id,typeFood)
		var allmenu []Menu
		if err != nil {
			panic(err.Error())
		}
<<<<<<< HEAD
		defer rows.Close()	
=======
		defer rows.Close()
				
>>>>>>> main
		for rows.Next() {
			var R_id int
			var M_id int
			var M_name string
			var M_price int
			var M_image_url string
			var M_type string

			err := rows.Scan(&M_id, &R_id, &M_name, &M_price, &M_type, &M_image_url)

			if err != nil {
				panic(err.Error())
			}
			menu := Menu{
				R_id:        R_id,
				M_id:        M_id,
				M_name:      M_name,
				M_price:     M_price,
				M_image_url: M_image_url,
				M_Extra: extraMenu(M_id),
			}
			
			allmenu = append(allmenu,menu)			
		}
		typ := Typ {
			T_name : typeFood,
			T_type : allmenu,

		}
		alltyp = append(alltyp, typ)
		return alltyp

}
func AllMenu(R_id int) []Typ {
	var alltyp []Typ
	rowst, err :=connectSqlDB().Query("select M_type from menu where R_id = ? GROUP BY M_type ",R_id)
	if err != nil {
		panic(err.Error())
	}
	defer rowst.Close()
	for rowst.Next(){
		var t string
		err := rowst.Scan(&t)

		if err != nil {
			panic(err.Error())
		}
		fmt.Println(t)
		alltyp = addData(alltyp,t,R_id)		

	}
	fmt.Println(alltyp)
	
	

	return alltyp

}

