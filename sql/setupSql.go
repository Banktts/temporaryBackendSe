package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("Go MYSQL Tutorial")
	db, errDb := sql.Open("mysql", "root:password@tcp(localhost:8080)/AgilestRelationDB")
	if errDb != nil {
		panic(errDb.Error())
	}

	CreatesFile, err := ioutil.ReadFile("sql/Table.sql")
	if err != nil {
		panic(err)
	}
	Craetes := strings.Split(string(CreatesFile), ";")
	for _, Craete := range Craetes {
		fmt.Println(Craete)
		if Craete != "" {
			if _, errI := db.Exec(Craete); errI != nil {
				panic(errI)
			}
		}

	}
	mockDatas, err := ioutil.ReadFile("sql/mockData.sql")
	if err != nil {
		panic(err)
	}
	Datas := strings.Split(string(mockDatas), ";")
	for _, Data := range Datas {
		fmt.Println(Data)
		if Data != "" {
			if _, errI := db.Exec(Data); errI != nil {
				panic(errI)
			}
		}

	}

	defer db.Close()
	fmt.Println("Successfully Connected to my Sql Db")
}
