package repository

import (
	"context"
	"time"

	"github.com/salamanderman234/peripheral-api/config"
	"github.com/salamanderman234/peripheral-api/domain"
	model "github.com/salamanderman234/peripheral-api/models"
	utility "github.com/salamanderman234/peripheral-api/utility"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type switchRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewSwitchRepository(connection *mongo.Client) domain.SwitchRepository {
	collection := connection.Database(config.GetDatabaseName()).Collection(config.SwitchsCollection)
	return &switchRepository{
		client:     connection,
		collection: collection,
	}
}

func (s *switchRepository) InsertSwitch(ctx context.Context, newSwitch model.Switch) error {
	_, err := s.collection.InsertOne(ctx, newSwitch)
	if err != nil {
		utility.NewLogEntry(nil).Error("Failed to insert data to switch collection in switch repository")
		return err
	}
	return nil
}

func (s *switchRepository) BatchInsertSwitchs(ctx context.Context, switchs []model.Switch) error {
	var switchsInterface []interface{}
	for _, element := range switchs {
		now := time.Now().Format(time.RFC1123)
		element.CreatedAt = now
		element.UpdateAt = now
		switchsInterface = append(switchsInterface, element)
	}
	_, err := s.collection.InsertMany(ctx, switchsInterface)
	if err != nil {
		utility.NewLogEntry(nil).Error("Failed to insert data to switch collection in switch repository")
		return err
	}
	return nil
}

func (s *switchRepository) UpdateSwitch(ctx context.Context, updateField model.Switch, filter model.Switch) error {

	return nil
}

func (s *switchRepository) DeleteSwitch(ctx context.Context, condition model.Switch) error {
	return nil
}

func (s *switchRepository) FindAllSwitchWithFilter(ctx context.Context, switchType string, switchManufacturer string, acforce float64, slug string) ([]model.Switch, error) {

	var switchs []model.Switch

	// making filter
	filter := bson.D{}
	if slug != "" {
		filter = append(filter, primitive.E{Key: "slug", Value: slug})
	}
	if switchType != "" {
		filter = append(filter, primitive.E{Key: "type", Value: switchType})
	}
	if switchManufacturer != "" {
		filter = append(filter, primitive.E{Key: "manufacturer", Value: switchManufacturer})
	}
	if acforce != 0.0 {
		filter = append(filter, primitive.E{Key: "actuation_force", Value: bson.D{
			primitive.E{Key: "$lte", Value: acforce},
		}})
	}

	// query
	cur, err := s.collection.Find(ctx, filter, nil)
	if err != nil {
		utility.NewLogEntry(nil).Error(err)
		return switchs, err
	}

	// to result
	if err = cur.All(ctx, &switchs); err != nil {
		utility.NewLogEntry(nil).Error(err)
		return switchs, err
	}

	return switchs, nil
}
