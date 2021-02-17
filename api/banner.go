package api
import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetBanner(lat float, long float) {	
	//var distance float
	distance := ((lat**2) + (long**2))**(1/2)
	db, errDb := sql.Open("mysql", "root:password@tcp(localhost:8080)/AgilestRelationDB")
	if errDb != nil {
		panic(errDb.Error())
	}

	rows, err:= db.Query("select R_ID, R_name from restaurant")
	if err!=nil{
		fmt.Println(err)
	}

	for rows.Next() {
		var R_ID int
		var R_name string
		err=rows.Scan(&R_ID, &R_name)
		fmt.Printf("id: %s name: %s\n", R_ID, R_name)
	}
}