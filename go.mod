module example.com

go 1.15

require (
	example.com/temporaryBackendSe/api v0.0.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gorilla/mux v1.8.0
	go.mongodb.org/mongo-driver v1.4.6
)

replace example.com/temporaryBackendSe/api => ../temporaryBackendSe/api
