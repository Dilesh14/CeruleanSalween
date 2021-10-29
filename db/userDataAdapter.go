package dataadapter

import (
	"context"
	"log"
	"time"

	services "github.com/dilesh14/CeruleanSalween/services"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ExternalId   string
	UserName     string
	EmailAddress string
	PhoneNumber  string
	Password     string
}

var connectionString string

func SetDbConnectionString(cs string) {
	connectionString = cs
}

func CreateUser(user User) bool {
	user.ExternalId = uuid.NewString()

	passHash, err := services.HashPassword(user.Password)

	if err != nil {
		log.Fatal("error in hashing password")
	}

	user.Password = passHash

	//connection to db
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("failed connection to client")
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	coll := client.Database("passManager").Collection("Users")
	proxy := client.Database("passManager").Collection("Proxy_Users")

	pResult, pErr := proxy.InsertOne(ctx, user)

	if pErr != nil {
		return false
	}

	log.Println("inserted the user to proxy:", pResult.InsertedID)

	result, inserErr := coll.InsertOne(ctx, user)
	if inserErr != nil {
		log.Fatal(inserErr)
	}

	log.Println("inserted new user, new userId: ", result.InsertedID)
	return true
}

func DeleteUser(id string) {

}

func AllUserSite(userId string) {

}
