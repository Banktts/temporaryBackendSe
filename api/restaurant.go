package api
import(
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
type Res struct {
	R_ID int `json:"R_ID"`
	R_name string `json:"R_name"`
	R_rating float64 `json:"R_rating"`
	R_votes int `json:"R_votes"`
	R_isRecomend bool `json:"R_isRecomend"`
	R_image_url string `json:"R_image_url"` 
	R_distance float64 `json:"R_distance"`
  }


func GetRestaurantBanner(keyword string, latitude float64, longitude float64)[]Res {	
	db := connectSqlDB()

	rows, err:= db.Query("select R_ID, R_name, R_rating, R_votes, R_isRecomend, R_image_url,sqrt(power((R_latitude - ?),2)+power((R_longitude - ?),2)) as R_distance from restaurant where R_name like '%"+keyword+"%' order by R_distance", latitude, longitude)
	if err!=nil{
		fmt.Println(err)
	}

	defer rows.Close()

	var restaurantList []Res
	for rows.Next() {
		var res Res
		err := rows.Scan(&res.R_ID, &res.R_name, &res.R_rating, &res.R_votes, &res.R_isRecomend, &res.R_image_url, &res.R_distance)
		if err != nil {
			panic(err.Error())
		}
		restaurantList = append(restaurantList, res)
	}
	return restaurantList
}