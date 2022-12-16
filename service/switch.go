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
	repository domain.SwitchRepository
}

func NewSwitchService(repo domain.SwitchRepository) domain.SwitchService {
	return &switchService{
		repository: repo,
	}
}

func (s *switchService) GetSwitch(ctx context.Context, filter entity.Switch, sort string) ([]entity.Switch, error) {
	// converting entity to model
	var filterModel model.Switch
	temp, _ := json.Marshal(filter)
	json.Unmarshal(temp, &filterModel)
	// calling repo
	switches, err := s.repository.FindAllSwitchWithFilter(ctx, filterModel, sort)
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
	// convert entity to model
	temp, _ := json.Marshal(switches)
	json.Unmarshal(temp, &switchesModel)
	// making slug for every switch
	for i := 0; i < len(switches); i++ {
		switchesModel[i].SwitchID = strings.Join(strings.Split(strings.ToLower(switches[i].Name), " "), "-")
	}
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

func (s *switchService) UpdateSwitch(ctx context.Context, updateField entity.Switch, filter entity.Switch) (int64, error) {
	// init
	var updateFieldModel model.Switch
	var filterModel model.Switch
	// making sure slug is empty
	if updateField.SwitchID != "" {
		updateField.SwitchID = ""
	}
	// creating new slug if there any new name
	if updateField.Name != "" {
		updateField.SwitchID = strings.Join(strings.Split(strings.ToLower(updateField.Name), " "), "-")
	}
	// convert to model
	temp, _ := json.Marshal(updateField)
	json.Unmarshal(temp, &updateFieldModel)
	// karena json encode di etity diset tidask ada maka harus dilakukan secara manual
	temp, _ = json.Marshal(filter)
	json.Unmarshal(temp, &filterModel)
	// calling repo
	modifiedDocuments, err := s.repository.UpdateSwitch(ctx, updateFieldModel, filterModel)
	if err != nil {
		return 0, err
	}
	return modifiedDocuments, nil
}

func (s *switchService) DeleteSwitch(ctx context.Context, filter string) (int64, error) {
	deletedCount, err := s.repository.DeleteSwitch(ctx, filter)
	if err != nil {
		return deletedCount, err
	}
	return deletedCount, nil
}

func (s *switchService) CountSwitch(ctx context.Context, filter entity.Switch) (int64, error) {
	// converting entity to model
	var filterModel model.Switch
	temp, _ := json.Marshal(filter)
	json.Unmarshal(temp, &filterModel)

	result, err := s.repository.CountSwitchWithFilter(ctx, filterModel)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *switchService) FindSimilarSwitch(ctx context.Context, switchEntity entity.Switch) ([]entity.Switch, error) {
	// making filter
	filterModel := model.Switch{
		Name: switchEntity.Name,
	}
	// find switch
	result, err := s.repository.FindAllSwitchWithFilter(ctx, filterModel, "")
	if err != nil {
		return nil, err
	}
	// decode to entity
	var entitiesResult []entity.Switch
	temp, _ := json.Marshal(result)
	json.Unmarshal(temp, &entitiesResult)
	// removing same switch as filter
	for index, element := range entitiesResult {
		if element == switchEntity {
			entitiesResult = append(entitiesResult[:index], entitiesResult[index+1:]...)
			break
		}
	}
	return entitiesResult, nil
}
