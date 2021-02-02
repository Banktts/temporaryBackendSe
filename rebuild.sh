docker-compose down
docker volume rm backend_BackendDb
docker-compose up -d
sleep 20
echo "rebuild data base complete"
go run setupSql.go



