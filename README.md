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

