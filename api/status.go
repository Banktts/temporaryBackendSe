package api
import(
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Deliveryman struct {
	D_ID int `json:"D_ID"`
	D_name string `json:"D_name"`
	D_rating float64 `json:"D_rating"`
	D_tel string `json:"D_phone"`
	Waiting_time string `json:"Waiting_time"`
  }

func OrderInfo(DeliveryId int, OrderId int)Deliveryman{
	db := connectSqlDB()

	var d Deliveryman
	err:= db.QueryRow("select D_ID, D_name, D_rating, D_phone from delivery_man where D_ID = ?" , DeliveryId).Scan(&d.D_ID, &d.D_name, &d.D_rating, &d.D_tel)
	if err!=nil{
		fmt.Println(err)
	}
	d.Waiting_time = "123"
	return d
}


 


