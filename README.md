GO MONGODB REACT
----------------

- [More about Go](https://go.dev/solutions/cloud)
- [Gin](https://pkg.go.dev/github.com/gin-gonic/gin) & [go Gin doc](https://go.dev/doc/tutorial/web-service-gin)


# Build a docker file for our Go API

[docker image](https://hub.docker.com/_/golang)
[See app Dockerfile](app/Dockerfile)

1. Create a module in which you can manage dependencies.

    go mod init demo/app
    go get .

2. Try it out locally

    go run .

    curl http://localhost:8080/ping

3. Build the image and test again

    docker build -t demo/go .
    docker image ls | grep demo
    docker run --rm -it --name gogo -p 8080:8080 -d demo/go
    docker ps
    docker logs <CONTAINER_ID>

    curl http://localhost:8080/ping

# Docker compose the Go app

[See latest references](https://docs.docker.com/compose/compose-file/03-compose-file/)

1. Run docker-compose

    docker-compose --profile test up -d
    docker ps
    docker-compose --profile test logs

2. Bring an env variable

Add this to [compose file](compose.yaml) :
    GIN_MODE=release

Then run :

    docker-compose --profile test up -d
    docker compose --profile test config
    docker-compose --profile test logs

Debug logs should be empty. Pings will show.


# Add MongoDB & mongo Express to docker compose

[See Bitnami image](https://hub.docker.com/r/bitnami/mongodb)
[Mongo Express](https://hub.docker.com/_/mongo-express)

We also define a bridge network and attache the services to it (app-tier).


    docker-compose --profile test up -d
    docker compose --profile test logs

    docker ps


# TROUBLESHOOTING

- Health check not working :

    docker exec -it 2b8f52351bed sh

    1. check if curl is installed
    2. check if wget is installed

    if 1 of them is then use it.

    docker-compose --profile test rm
    docker-compose --profile test up -d


- Express not running

    docker ps -a | grep express
    docker logs 863dac80c19b

    -> fix url

    + Authentication failed.
    -> fix user access