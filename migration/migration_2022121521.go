package main

import (
	"context"

	"github.com/salamanderman234/peripheral-api/config"
	"github.com/salamanderman234/peripheral-api/utility"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// membaca config json
func init() {
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()

	if err != nil {
		utility.NewLogEntry(nil).Panic(err)
	}

	logrus.SetFormatter(&prefixed.TextFormatter{})
}

func main() {
	// make connection to database
	client, err := config.ConnectDB(config.GetDatabaseUri())
	if err != nil {
		utility.NewLogEntry(nil).Panic(err)
	}

	// create index
	_, err = client.Database(config.GetDatabaseName()).Collection(config.SwitchesCollection).Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    map[string]int{"email": 1},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		panic(err)
	}
}
