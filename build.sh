GOARCH=amd64 GOOS=linux go build -o timely

docker build -t timely .

docker run -d timely
echo timely server running as a docker container
echo ------------------------------------------
echo 'type docker ps to view container'
