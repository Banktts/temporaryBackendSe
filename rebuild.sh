docker-compose down
docker volume rm backend_BackendDb
docker-compose up -d
sleep 20
sh dbBuild.sh
echo "rebuild data base complete"




