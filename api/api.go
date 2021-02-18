package api
import(
	"context"
	"database/sql"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "github.com/go-sql-driver/mysql"
)
var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
var client, errClient = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
var sqlDb, errDb = sql.Open("mysql", "root:password@tcp(localhost:8080)/AgilestRelationDB")

func connectMongoDB() *mongo.Database {

	if errClient != nil {
		fmt.Println(errClient)
	}
	//Call the connect function of client
	errorCon := client.Connect(ctx)
	if errorCon != nil {
		fmt.Println(errorCon)
	}
	// Check the connection
	errClient = client.Ping(context.TODO(), nil)

	if errClient != nil {
		fmt.Println(errClient)
	}

	fmt.Println("Connected to MongoDB!")
	return client.Database("AgilestNoRelationDB")

}

func disconectMongoDB(){
	defer cancel()
	fmt.Println("Connection to MongoDB closed.")
}

func connectSqlDB() *sql.DB {
	if errDb != nil {
		panic(errDb.Error())
	}
	return sqlDb
}

func disconnectSqlDB(){
	defer sqlDb.Close()
	fmt.Println("Successfully Connected to my Sql Db")
}
