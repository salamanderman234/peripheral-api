package repository

import (
	"context"
	"errors"
	"time"

	"github.com/salamanderman234/peripheral-api/config"
	"github.com/salamanderman234/peripheral-api/domain"
	model "github.com/salamanderman234/peripheral-api/models"
	"github.com/salamanderman234/peripheral-api/utility"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type authRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewAuthRepository(client *mongo.Client) domain.AuthRepository {
	collection := client.Database(config.GetDatabaseName()).Collection(config.AuthCollection)
	return &authRepository{
		client:     client,
		collection: collection,
	}
}

func (a *authRepository) CreateNewUser(ctx context.Context, user model.User) (model.User, error) {
	now := time.Now().Format(time.RFC1123)
	user.CreatedAt = now
	user.UpdateAt = now
	_, err := a.collection.InsertOne(ctx, user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
func (a *authRepository) GetUserByCredentials(ctx context.Context, username string, password string) (model.User, error) {
	var resultModel model.User
	filter := bson.M{"username": username, "password": password}
	err := a.collection.FindOne(ctx, filter).Decode(&resultModel)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		go utility.NewLogEntry(nil).Error(err)
		return model.User{}, err
	}
	return resultModel, nil
}
