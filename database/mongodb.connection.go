package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() {
	if err := godotenv.Load(); err != nil {
		log.Println("No se encuentra el archivo .env")
	}
	uri := os.Getenv("MONGODB_URI")
	//fmt.Println(uri)

	if uri == "" {
		log.Fatal("Se necesitan datos en la uri para que esto funcione")
	}

	//creo el newClient y ejecuto la conexion pasando la uri mongoDb al ApplyUri
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	//funcion que se ejecuta al final cuando el client se desconecta
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	//el bson M es una representacion desordenada de un documento lo que seria un MAP y el bson D es la representacion ordenada de un documento lo que seria un SLICE
	var result bson.M

	//el metodo Decode devuelve un error segun su tipo si todo sale bien y puede clasificar y desclasificar el item que se requirio este devuelve un nil
	if err := client.Database("goCrud").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}

	fmt.Println("Conectado correctamente a la base de datos")

}
