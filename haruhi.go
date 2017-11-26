package main

import (
	"gopkg.in/mgo.v2"
)

var MongoSession *mgo.Session

func main() {
	MongoSession := CreateMongoSession("http://localhost")
	defer MongoSession.Close()

	CreateManagerServer(7777)
}
