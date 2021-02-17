package api
import(
	"context"
	"database/sql"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "github.com/go-sql-driver/mysql"
)
var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
var client, errClient = mongo.Connect(context.TODO(), clientOptions)
var sqlDb, errDb = sql.Open("mysql", "root:password@tcp(localhost:8080)/AgilestRelationDB")

func connectMongoDB() *mongo.Database {

	if errClient != nil {
		fmt.Println(errClient)
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
	errDis := client.Disconnect(context.TODO())
	if errDis != nil {
		fmt.Println(errClient)
	}
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
