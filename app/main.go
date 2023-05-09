package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "os"
  "fmt"
  "context"
	"log"
  "encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
  r := gin.Default()
  r.GET("/ping", health)
  r.GET("/students", students)
  r.Run() // listen and serve on 0.0.0.0:8080
}

func health(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "message": "pong",
  })
}

func students(c *gin.Context) {

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	log.Println("URIL : ", uri)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Test the connection
	err = client.Ping(c, nil)
	if err != nil {
	  log.Fatal(err)
	}

	defer func() {
		log.Println("About to disconnect !")
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("example").Collection("students")
	fname := "Ron"
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"fname", fname}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", fname)
		return
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
  c.IndentedJSON(http.StatusOK, result)
}