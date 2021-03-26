package api
import(
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Deliveryman struct {
	D_ID int `json:"R_ID"`
	D_name string `json:"R_name"`
	D_rating float64 `json:"R_rating"`
	D_tel int `json:"R_votes"`
  }

func Delivery()[]Res {	
	db := connectSqlDB()

	rows, err:= db.Query("select R_ID, R_name, R_rating, R_votes, R_isRecomend, R_image_url, ROUND((sqrt(power((R_latitude-?)*110.574,2)+power((R_longitude-?)*111.320,2))),2) as R_distance, R_location from restaurant where R_name like '%"+keyword+"%' order by R_distance", latitude, longitude)
	if err!=nil{
		fmt.Println(err)
	}

	defer rows.Close()

	var restaurantList []Res
	for rows.Next() {
		
		err := rows.Scan(&res.R_ID, &res.R_name, &res.R_rating, &res.R_votes, &res.R_isRecomend, &res.R_image_url, &res.R_distance, &res.R_location)
		if err != nil {
			panic(err.Error())
		}
		
	}
	return restaurantList
}
func WaitingTime() int{
	db := connectSqlDB()



	rows, err:= db.Query("select D_ID from delivery_man ",D_ID)
	if err!=nil{
		fmt.Println(err)
	}


	rows, err:= db.Query("select ROUND((sqrt(power((R_latitude-?)*110.574,2)+power((R_longitude-?)*111.320,2))),2) as R_distance, R_location from restaurant where R_name like '%"+keyword+"%' order by R_distance", latitude, longitude)
	if err!=nil{
		fmt.Println(err)
	}
	defer rows.Close()

	var restaurantList []Res
	for rows.Next() {
		
		err := rows.Scan()
		if err != nil {
			panic(err.Error())
		}
		
	}
	return 

}

func OrderInfo() {
	
}


 


