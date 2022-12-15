package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Conn struct {
}

const (
	mongoUrl = "mongodb://192.168.1.9:27017"
	auth     = "admin"
	username = "admin"
	password = "admin"
)

var client *mongo.Client

func (c *Conn) conn() {
	cred := options.Credential{AuthSource: auth, Username: username, Password: password}
	clientOptions := options.Client().ApplyURI(mongoUrl).SetAuth(cred)

	con, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	} else {
		log.Println("connect mongodb success:", clientOptions.GetURI())
	}
	client = con
}

func (c *Conn) Close() {
	if client != nil {
		client.Disconnect(nil)
	}
}

func (c *Conn) GetCol(dbName string, colName string) mongo.Collection {
	if client == nil {
		c.conn()
	}
	col := client.Database(dbName).Collection(colName)
	log.Println("mongodb collection:", c.GetColFullName(*col))
	return *col
}

func (c *Conn) GetColFullName(coll mongo.Collection) string {
	return coll.Database().Name() + "." + coll.Name()
}
