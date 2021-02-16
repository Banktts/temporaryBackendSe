echo "db build"
go get github.com/go-sql-driver/mysql
go get go.mongodb.org/mongo-driver
echo "go build succeed"
go run sql/setupSql.go
go run mongo/setupMongo.go