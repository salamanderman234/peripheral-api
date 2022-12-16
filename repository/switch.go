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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type switchRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewSwitchRepository(connection *mongo.Client) domain.SwitchRepository {
	collection := connection.Database(config.GetDatabaseName()).Collection(config.SwitchesCollection)
	return &switchRepository{
		client:     connection,
		collection: collection,
	}
}

func (s *switchRepository) InsertSwitch(ctx context.Context, newSwitch model.Switch) error {
	_, err := s.collection.InsertOne(ctx, newSwitch)
	if err != nil {
		go utility.NewLogEntry(nil).Error(err)
		return err
	}
	return nil
}

func (s *switchRepository) BatchInsertSwitches(ctx context.Context, switches []model.Switch) ([]interface{}, error) {
	var switchesInterface []interface{}
	// set updateat and convert into []interface
	for _, element := range switches {
		now := time.Now().Format(time.RFC1123)
		element.UpdateAt = now
		switchesInterface = append(switchesInterface, element)
	}
	// query
	result, err := s.collection.InsertMany(ctx, switchesInterface)
	if err != nil {
		go utility.NewLogEntry(nil).Error(err)
		return result.InsertedIDs, err
	}
	return result.InsertedIDs, nil
}

func (s *switchRepository) UpdateSwitch(ctx context.Context, updateField model.Switch, filter primitive.M) (int64, error) {
	// convert into bson
	filterBson := bson.M{
		"$and": []primitive.M{filter},
	}
	// set updateat
	now := time.Now().Format(time.RFC1123)
	updateField.UpdateAt = now
	// convert into bson for update field
	updateFieldBson := bson.M{
		"$set": updateField,
	}

	// query
	result, err := s.collection.UpdateMany(ctx, filterBson, updateFieldBson, options.Update())

	if err != nil {
		go utility.NewLogEntry(nil).Error(err)
		return 0, err
	}
	return result.ModifiedCount, nil
}

func (s *switchRepository) DeleteSwitch(ctx context.Context, condition model.Switch) error {
	return nil
}

func (s *switchRepository) FindAllSwitchWithFilter(ctx context.Context, filter model.Switch) ([]model.Switch, error) {
	// init
	var switches []model.Switch

	// making filter
	filterBson := bson.M{}
	if filter.Slug != "" {
		filterBson["slug"] = filter.Slug
	}
	if filter.Type != "" {
		filterBson["type"] = filter.Type
	}
	if filter.Manufacturer != "" {
		filterBson["manufacturer"] = filter.Manufacturer
	}
	if filter.ActuationForce != 0.0 {
		filterBson["actuation_force"] = filter.ActuationForce
	}

	// query
	cur, err := s.collection.Find(ctx, filterBson, nil)
	if err != nil {
		go utility.NewLogEntry(nil).Error(err)
		return switches, err
	}

	// to result
	if err = cur.All(ctx, &switches); err != nil {
		go utility.NewLogEntry(nil).Error(err)
		return switches, err
	}

	return switches, nil
}
