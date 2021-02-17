package api
import(
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
type Res struct {
	R_ID int `json:"id"`
	R_name string `json:"R_name"`
  }


func GetRestaurant(id int)[]Res {	
	db := connectSqlDB()

	rows, err:= db.Query("select R_ID, R_name from restaurant WHERE R_ID = ?", id)
	if err!=nil{
		fmt.Println(err)
	}

	defer rows.Close()

	var restaurantList []Res
	for rows.Next() {
		var res Res
		err := rows.Scan(&res.R_ID, &res.R_name)
		if err != nil {
			panic(err.Error())
		}
		restaurantList = append(restaurantList, res)
	}

	return restaurantList
}