package service

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/salamanderman234/peripheral-api/domain"
	"github.com/salamanderman234/peripheral-api/entity"
	model "github.com/salamanderman234/peripheral-api/models"
	"github.com/salamanderman234/peripheral-api/policy"
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

func (s *switchService) CreateSwitch(ctx context.Context, switchs []entity.Switch) (*[]*policy.SwitchPolicy, error) {
	var switchsModel []model.Switch
	var policyResult []*policy.SwitchPolicy

	// checking policy
	for i := 0; i < len(switchs); i++ {
		result := policy.InsertSwitchPolicy(switchs[i])
		if result != nil {
			policyResult = append(policyResult, result)
		} else {
			// make slug
			switchs[i].Slug = strings.Join(strings.Split(strings.ToLower(switchs[i].Name), " "), "-")
		}
	}

	if len(policyResult) != 0 {
		return &policyResult, nil
	}

	// convert entity to model
	jsonSwitchs, _ := json.Marshal(switchs)
	json.Unmarshal(jsonSwitchs, &switchsModel)

	// calling repo
	err := s.switchRepo.BatchInsertSwitchs(ctx, switchsModel)
	if err != nil {
		return nil, err
	}

	return nil, nil
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

func (s *switchService) UpdateSwitch(ctx context.Context, updateField entity.Switch, filter entity.Switch) ([]byte, error) {
	return nil, nil
}
