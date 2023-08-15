# Run proxy-golang-api container in detached mode
docker run -d --rm -p 8080:8080 --name proxy-golang-api proxy-golang-api

# Run proxy-golang-psql container in detached mode
docker run -d --rm -p 5432:5432 --name proxy-golang-psql proxy-golang-psql
