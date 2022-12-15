package service

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/salamanderman234/peripheral-api/domain"
	"github.com/salamanderman234/peripheral-api/entity"
	model "github.com/salamanderman234/peripheral-api/models"
)

type switchService struct {
	switchRepo domain.SwitchRepository
}

func NewSwitchService(repo domain.SwitchRepository) domain.SwitchService {
	return &switchService{
		switchRepo: repo,
	}
}

func (s *switchService) GetSwitch(ctx context.Context, filter entity.Switch) ([]byte, error) {
	switchs, err := s.switchRepo.FindAllSwitchWithFilter(ctx, filter.Type, filter.Manufacturer, filter.ActuationForce, filter.Slug)
	if err != nil {
		return nil, err
	}

	switchsParshed, _ := json.Marshal(switchs)
	return switchsParshed, nil
}

func (s *switchService) CreateSwitch(ctx context.Context, switchs []entity.Switch) ([]interface{}, error) {
	var switchsModel []model.Switch

	// checking policy
	for i := 0; i < len(switchs); i++ {
		switchs[i].Slug = strings.Join(strings.Split(strings.ToLower(switchs[i].Name), " "), "-")
	}

	// convert entity to model
	jsonSwitchs, _ := json.Marshal(switchs)
	json.Unmarshal(jsonSwitchs, &switchsModel)

	// calling repo
	insertedId, err := s.switchRepo.BatchInsertSwitchs(ctx, switchsModel)
	if err != nil {
		return nil, err
	}

	return insertedId, nil
}

func (s *switchService) CreateOneSwitch(ctx context.Context, switchEntity entity.Switch) error {
	var switchModel model.Switch
	// convert entity to model
	jsonSwitchs, _ := json.Marshal(switchEntity)
	json.Unmarshal(jsonSwitchs, &switchModel)

	// calling repo
	err := s.switchRepo.InsertSwitch(ctx, switchModel)
	if err != nil {
		return err
	}
	return nil
}

func (s *switchService) UpdateSwitch(ctx context.Context, updateField entity.Switch, filter entity.Switch) (int64, error) {
	var updateFieldModel model.Switch
	var filterModel model.Switch

	// creating new slug if there any new name
	if updateField.Name != "" {
		updateField.Slug = strings.Join(strings.Split(strings.ToLower(updateField.Name), " "), "-")
	}
	// convert to model
	temp, _ := json.Marshal(updateField)
	json.Unmarshal(temp, &updateFieldModel)
	temp, _ = json.Marshal(filter)
	json.Unmarshal(temp, &filterModel)

	// calling repo
	modifiedDocument, err := s.switchRepo.UpdateSwitch(ctx, updateFieldModel, filterModel)
	if err != nil {
		return 0, err
	}
	return modifiedDocument, nil
}
