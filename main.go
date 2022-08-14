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
		fmt.Println("Erro de conexão com Mongo: ", err)
		os.Exit(1)
	}

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	collection := client.Database("MeuDb").Collection("Filmes")
	fmt.Println("Collection Type: ", reflect.TypeOf(collection))

	oneDoc := Filme{
		Nome:       "Rocky",
		Nota:       10,
		Lancamento: false,
	}

	fmt.Println("oneDoc Type: ", reflect.TypeOf(oneDoc))

	result, insertErr := collection.InsertOne(ctx, oneDoc)
	if insertErr != nil {
		fmt.Println("InsertONE Error:", insertErr)
		os.Exit(1)
	} else {
		fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
		fmt.Println("InsertOne() api result type: ", result)

		newID := result.InsertedID
		fmt.Println("InsertedOne(), newID", newID)
		fmt.Println("InsertedOne(), newID type:", reflect.TypeOf(newID))
	}

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Leitura Error:", err)
		os.Exit(1)
	}

	var dados []bson.M
	if err = cursor.All(ctx, &dados); err != nil {
		fmt.Println("Leitura de dados Error:", err)
	}
	fmt.Println(dados)
}
