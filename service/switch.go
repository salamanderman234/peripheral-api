package service

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/salamanderman234/peripheral-api/domain"
	"github.com/salamanderman234/peripheral-api/entity"
	model "github.com/salamanderman234/peripheral-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type switchService struct {
	repository domain.SwitchRepository
}

func NewSwitchService(repo domain.SwitchRepository) domain.SwitchService {
	return &switchService{
		repository: repo,
	}
}

func (s *switchService) GetSwitch(ctx context.Context, filter entity.Switch) ([]entity.Switch, error) {
	// converting entity to model
	var filterModel model.Switch
	temp, _ := json.Marshal(filter)
	json.Unmarshal(temp, &filterModel)
	if filter.Slug != "" {
		filterModel.Slug = filter.Slug
	}
	// calling repo
	switches, err := s.repository.FindAllSwitchWithFilter(ctx, filterModel)
	if err != nil {
		return nil, err
	}
	// parsing result into json format so controller can convert it back into entity
	var entitiesSwitches []entity.Switch
	temp, _ = json.Marshal(switches)
	json.Unmarshal(temp, &entitiesSwitches)
	return entitiesSwitches, nil
}

func (s *switchService) CreateSwitch(ctx context.Context, switches []entity.Switch) ([]interface{}, error) {
	// init
	var switchesModel []model.Switch
	// making slug for every switch
	for i := 0; i < len(switches); i++ {
		switches[i].Slug = strings.Join(strings.Split(strings.ToLower(switches[i].Name), " "), "-")
	}
	// convert entity to model
	temp, _ := json.Marshal(switches)
	json.Unmarshal(temp, &switchesModel)
	// calling repo
	insertedId, err := s.repository.BatchInsertSwitches(ctx, switchesModel)
	if err != nil {
		return nil, err
	}

	return insertedId, nil
}

func (s *switchService) CreateOneSwitch(ctx context.Context, switchEntity entity.Switch) error {
	// init
	var switchModel model.Switch
	// convert entity to model
	temp, _ := json.Marshal(switchEntity)
	json.Unmarshal(temp, &switchModel)
	// calling repo
	err := s.repository.InsertSwitch(ctx, switchModel)
	if err != nil {
		return err
	}
	return nil
}

// ganti filter menjadi map agar bisa diencode menjadi bson
func (s *switchService) UpdateSwitch(ctx context.Context, updateField entity.Switch, filter entity.Switch) (int64, error) {
	// init
	var updateFieldModel model.Switch
	var filterModel primitive.M
	// creating new slug if there any new name
	if updateField.Name != "" {
		updateField.Slug = strings.Join(strings.Split(strings.ToLower(updateField.Name), " "), "-")
	}
	// making sure slug is empty
	if updateField.Slug != "" {
		updateField.Slug = ""
	}
	// convert to model
	temp, _ := json.Marshal(updateField)
	json.Unmarshal(temp, &updateFieldModel)
	// karena json encode di etity diset tidask ada maka harus dilakukan secara manual
	temp, _ = json.Marshal(filter)
	json.Unmarshal(temp, &filterModel)
	if filter.Slug != "" {
		filterModel["slug"] = filter.Slug
	}
	// calling repo
	modifiedDocuments, err := s.repository.UpdateSwitch(ctx, updateFieldModel, filterModel)
	if err != nil {
		return 0, err
	}
	return modifiedDocuments, nil
}
