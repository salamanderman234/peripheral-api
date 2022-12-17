package repository

import (
	"context"

	"github.com/salamanderman234/peripheral-api/config"
	"github.com/salamanderman234/peripheral-api/domain"
	model "github.com/salamanderman234/peripheral-api/models"
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
	_, err := a.collection.InsertOne(ctx, user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
func (a *authRepository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	var resultModel model.User
	filter := bson.M{"username": username}
	err := a.collection.FindOne(ctx, filter).Decode(&resultModel)
	if err != nil {
		return model.User{}, err
	}
	return resultModel, nil
}
