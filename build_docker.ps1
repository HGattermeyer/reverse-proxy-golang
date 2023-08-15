cd database
docker build -t proxy-golang-psql .

cd ../proxy-golang
docker build -t proxy-golang-api .
