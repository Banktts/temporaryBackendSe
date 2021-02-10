docker-compose down
path=`pwd`
dirname=`basename $path | awk '{print tolower($0)}'`
docker volume rm ${dirname}_BackendDb
docker-compose up -d
sleep 20
sh dbBuild.sh
echo "rebuild data base complete"




