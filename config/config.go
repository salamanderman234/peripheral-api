package config

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	SwitchesCollection = "switches"
	KeyboardCollection = "keyboards"
	MouseCollection    = "mouses"
)

func GetDatabaseUri() string {
	driver := viper.GetString("database.driver")
	// username := viper.GetString("database.username")
	// pass := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	return fmt.Sprintf("%s://%s:%s/?maxPoolSize=20&w=majority", driver, host, port)
}

func CreateAppPath(path string) string {
	appUrl := viper.GetString("app.url")
	appVersion := viper.GetString("app.version")

	return fmt.Sprintf("%s/v%s%s", appUrl, appVersion, path)
}

func GetDatabaseName() string {
	return viper.GetString("database.name")
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
