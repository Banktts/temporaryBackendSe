package api
import(
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetBanner(latitude float64, longitude float64)[]Res {	
	db := connectSqlDB()

	rows, err:= db.Query("select R_ID, R_name, R_rating, R_votes, R_isRecomend, R_image_url, (sqrt(power((R_latitude-?),2)+power((R_longitude-?),2)))*100 as R_distance from restaurant where R_isRecomend = 1 order by R_distance limit 5" , latitude, longitude)
	if err!=nil{
		fmt.Println(err)
	}

	defer rows.Close()

	var bannerList []Res
	for rows.Next() {
		var res Res
		err := rows.Scan(&res.R_ID, &res.R_name, &res.R_rating, &res.R_votes, &res.R_isRecomend, &res.R_image_url, &res.R_distance)
		if err != nil {
			panic(err.Error())
		}
		bannerList = append(bannerList, res)
	}
	return bannerList
}