GO MONGODB REACT
----------------

- [More about Go](https://go.dev/solutions/cloud)
- [Gin](https://pkg.go.dev/github.com/gin-gonic/gin) & [go Gin doc](https://go.dev/doc/tutorial/web-service-gin)


# Build a docker file for our Go API

[docker image](https://hub.docker.com/_/golang)
[See app Dockerfile](app/Dockerfile)

1. Create a module in which you can manage dependencies.

    export PATH=$PATH:/opt/go/bin
    go mod init demo/app
    go get .

2. Try it out locally

    go build
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


# Provide health check on mongodb and dependency for express

```yaml
    healthcheck:
      test: echo 'db.runCommand("ping").ok' | mongosh localhost:27017/test --quiet
      interval: 10s
      timeout: 10s
      retries: 3
      start_period: 20s
```
# Prepare the database

- Create a db : example 

- Add a user with credentials
> db.createUser(
  {
    user: "test",
    pwd:  "password123",   
    roles: [ { role: "readWrite", db: "example" } ],
    passwordDigestor : "server"
  }
)

- Prepare the data :

> db.createCollection("students")
> db.students.insertOne({
            "fname":"Ron", 
            "city":"United States of America", 
            "courses":[
                         "python", 
                         "django", 
                         "node"
                      ]
})

> db.students.find()

# Add an endpoint in the go application

Add the following depedencies to communicate with MongoDB.

```go
"os"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
```

And do not forget about the environment variable in the docker compose file :
[See .env file](.env.example)

- Rebuild the image :

    docker-compose --profile test build gogin

# Add VueJs/React/Angular frontend



# Going further

- Bring authentication
- Bring Istio


---------------
TROUBLESHOOTING
---------------


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

    1. Expose mongodb port
    2. access mongo from local mongo shell and make sure the password is working for root

        mongo --authenticationDatabase "admin" -u root -p
        or
        mongosh --port 27017 -u root -p 'example' --authenticationDatabase 'admin'

        mongosh test_database --port 27017 -u test_user -p 'password123' --authenticationDatabase 'admin'
        mongosh test_database --port 27017 -u test_user -p 'password123' --authenticationDatabase 'example'

- Clean up

    3. Clean up everything

        docker-compose down -v --rmi all

    => not enough -> the volume was actually binded to host and not tight to docker volume
    docker volumes ls is not showing up that kind of volume...
    See inspect the container

    -> Solution create a volume (or/and add the volumes: line to docker compose)
    -> Or/and remove the host data

- MongoDB connection issue

    [Some doc](https://mongodb.github.io/mongo-java-driver/3.8/javadoc/com/mongodb/ConnectionString.html)
    [Other doc about driver](https://www.digitalocean.com/community/tutorials/how-to-use-go-with-mongodb-using-the-mongodb-go-driver)

- Change the password of a user :
    
    > db.changeUserPassword(username, password)