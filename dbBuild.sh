echo "db build"
go get github.com/go-sql-driver/mysql
go clean -modcache  
go mod init example.com 
go mod tidy
echo "go build succeed"
go run setupSql.go