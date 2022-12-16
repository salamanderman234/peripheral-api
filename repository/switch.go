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
		element.CreatedAt = now
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

func (s *switchRepository) DeleteSwitch(ctx context.Context, condition string) (int64, error) {
	filter := bson.M{"slug": condition}
	result, err := s.collection.DeleteOne(ctx, filter)
	if err != nil {
		go utility.NewLogEntry(nil).Error(err)
		return 0, err
	}
	return result.DeletedCount, nil
}

func (s *switchRepository) FindAllSwitchWithFilter(ctx context.Context, filter model.Switch, sort string) ([]model.Switch, error) {
	// init
	var switches []model.Switch

	// making filter
	filterBson := bson.M{}
	if filter.SwitchID != "" {
		filterBson["switch_id"] = filter.SwitchID
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

	// makin sorting field
	defaultSort := "manufacturer"
	if sort != "" {
		defaultSort = sort
	}
	sortBson := bson.D{
		{Key: defaultSort, Value: 1},
		{Key: "name", Value: 1},
		{Key: "updateat", Value: -1},
	}
	opts := options.Find().SetSort(sortBson)

	// query
	cur, err := s.collection.Find(ctx, filterBson, opts)
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
