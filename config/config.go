package config

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDatabaseUri() string {
	driver := viper.GetString("database.driver")
	username := viper.GetString("database.username")
	pass := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	return fmt.Sprintf("%s://%s:%s@%s%s/?maxPoolSize=20&w=majority", driver, username, pass, host, port)
}

func ConnectDB(uri string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	return client, nil
}
