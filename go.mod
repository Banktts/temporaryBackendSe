module example.com

go 1.15

require (
	github.com/go-sql-driver/mysql v1.5.0
	go.mongodb.org/mongo-driver v1.4.6
	example.com/temporaryBackendSe/api v0.0.0
)

replace example.com/temporaryBackendSe/api => ../temporaryBackendSe/api