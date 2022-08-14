package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Filme struct {
	Nome       string `json:"Nome"`
	Nota       int    `json:"Nota"`
	Lancamento bool   `json:"Lancamento"`
}

func main() {
	uri := "sua_string_de_conexao"
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.Connect Error: ", err)
		os.Exit(1)
	}

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	collection := client.Database("MeuDb").Collection("Filmes")
	fmt.Println("Collection Type: ", reflect.TypeOf(collection))

	filmes := []interface{}{
		Filme{
			Nome:       "Rocky",
			Nota:       10,
			Lancamento: false,
		},
		Filme{
			Nome:       "Carter",
			Nota:       7,
			Lancamento: true,
		},
		Filme{
			Nome:       "Hulk",
			Nota:       6,
			Lancamento: false,
		},
	}

	resultMany, insertErr := collection.InsertMany(ctx, filmes)
	if insertErr != nil {
		fmt.Println("InsertMany Error:", insertErr)
		os.Exit(1)
	} else {
		fmt.Println("InsertMany(), newIDs", resultMany.InsertedIDs)
	}

	filme := Filme{
		Nome:       "Thor",
		Nota:       6,
		Lancamento: false,
	}

	resultOne, insertErr := collection.InsertOne(ctx, filme)
	if insertErr != nil {
		fmt.Println("InsertOne Error:", insertErr)
		os.Exit(1)
	} else {
		newID := resultOne.InsertedID
		fmt.Println("InsertedOne(), newID", newID)
	}

	// Consultar registros da collection
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Find Error:", err)
		os.Exit(1)
	}

	var data []bson.M
	err = cursor.All(ctx, &data)
	if err != nil {
		fmt.Println("Cursor All Error:", err)
	}
	fmt.Println(data)
}
