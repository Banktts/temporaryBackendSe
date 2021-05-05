package api
import(
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math"
)

type Deliveryman struct {
	D_ID int `json:"D_ID"`
	D_name string `json:"D_name"`
	D_rating float64 `json:"D_rating"`
	D_tel string `json:"D_phone"`
	Waiting_time int `json:"Waiting_time"`
  }

func OrderInfo(DeliveryId int, OrderId int)Deliveryman{
	db := connectSqlDB()

	var d Deliveryman
	err:= db.QueryRow("select D_ID, D_name, D_rating, D_phone from delivery_man where D_ID = ?" , DeliveryId).Scan(&d.D_ID, &d.D_name, &d.D_rating, &d.D_tel)
	if err!=nil{
		fmt.Println(err)
	}
	d.Waiting_time = int(WaitingTime(DeliveryId,OrderId))
	return d
}

func WaitingTime(D_ID int,O_ID int) float64{
    db := connectSqlDB()
    var D_latitude float64
    var D_longitude float64
    var C_latitude float64
    var C_longitude float64

    rows, err:= db.Query("select D_latitude,D_longitude from delivery_man where D_ID = ? ",D_ID)
    if err!=nil{
        fmt.Println(err)
    }
    for rows.Next() {
        
        err := rows.Scan(&D_latitude,&D_longitude)
        if err != nil {
            panic(err.Error())
        }
        
    }
    rows2, err2:= db.Query("select customer.C_latitude,customer.C_longitude from ordert natural join customer where O_ID = ? ",O_ID)
    if err2!=nil{
        fmt.Println(err2)
    }
    for rows2.Next() {
        
        err2 := rows2.Scan(&C_latitude,&C_longitude)
        if err2 != nil {
            panic(err2.Error())
        }
        
    }
    return 20+ math.Sqrt(math.Pow((D_latitude-C_latitude)*110.574,2) + math.Pow((D_longitude-C_longitude)*111.320,2))
}

 


