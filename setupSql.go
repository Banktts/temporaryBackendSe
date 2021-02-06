package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	fmt.Println("Go MYSQL Tutorial")
	db,errDb:=sql.Open("mysql","root:password@tcp(localhost:8080)/testDb")
	if errDb != nil{
		panic(errDb.Error())
	}
	
	
	insertFile, err:=ioutil.ReadFile("sql/insert_into_db_aroijangfood.sql")
	if err != nil{
		panic(err)
	}
	inserts := strings.Split(string(insertFile), ";")
	for _, insert := range inserts {
		fmt.Println(insert)
		if _,errI := db.Exec(insert); errI != nil{
			panic(errI)
		}
	}


	defer db.Close()
	fmt.Println("Successfully Connected to my Sql Db")
	
}
