package api
import(
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
var client, errClient = mongo.Connect(context.TODO(), clientOptions)

func Test(){
	fmt.Println("test")
}

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
