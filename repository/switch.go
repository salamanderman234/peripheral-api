package repository

import (
	"context"

	"github.com/salamanderman234/peripheral-api/config"
	"github.com/salamanderman234/peripheral-api/domain"
	model "github.com/salamanderman234/peripheral-api/models"
	utility "github.com/salamanderman234/peripheral-api/utility/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type switchRepository struct {
	client *mongo.Client
}

func NewSwitchRepository(connection *mongo.Client) domain.SwitchRepository {
	return &switchRepository{
		client: connection,
	}
}

func (s *switchRepository) InsertSwitch(ctx context.Context, newSwitch model.Switch) (model.Switch, error) {
	return model.Switch{}, nil
}

func (s *switchRepository) BatchInsertSwitchs(ctx context.Context, switchs []model.Switch) ([]model.Switch, error) {
	return []model.Switch{}, nil
}

func (s *switchRepository) UpdateSwitch(ctx context.Context, updateField model.Switch, condition model.Switch) error {
	return nil
}

func (s *switchRepository) DeleteSwitch(ctx context.Context, condition model.Switch) error {
	return nil
}

func (s *switchRepository) FindAllSwitchWithFilter(ctx context.Context, filter model.Switch) ([]model.Switch, error) {

	var switchs []model.Switch

	collection := s.client.Database(config.GetDatabaseName()).Collection(config.SwitchsCollection)
	filterBSON, _ := bson.Marshal(filter)

	cur, err := collection.Find(ctx, filterBSON, nil)
	if err != nil {
		utility.NewLogEntry(nil).Error("failed to get data from switch collection in switch repository")
		return switchs, err
	}

	if err = cur.All(ctx, &switchs); err != nil {
		utility.NewLogEntry(nil).Error("failed to decode data from cursor collections in switch repository")
		return switchs, err
	}

	return switchs, nil
}
