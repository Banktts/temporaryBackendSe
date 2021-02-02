package main

import(
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

func main(){
	fmt.Println("Go MYSQL Tutorial")
	db,err:=sql.Open("mysql","root:pasword@tcp(localhost:3306)/testDb5")
	if err != nil{
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Successfully Connected to my Sql Db")
}
