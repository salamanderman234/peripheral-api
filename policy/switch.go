package policy

import (
	"context"
	"strings"

	"github.com/salamanderman234/peripheral-api/domain"
	"github.com/salamanderman234/peripheral-api/entity"
)

type SwitchPolicy struct {
	Data                entity.Switch `json:"data"`
	NameMessage         string        `json:"name,omitempty"`
	ManufacturerMessage string        `json:"manufacturer,omitempty"`
	TypeMessage         string        `json:"type,omitempty"`
}

func DocumentSwitchPolicy(ctx context.Context, switcEntity entity.Switch, service domain.SwitchService, op string) *SwitchPolicy {
	var policy SwitchPolicy
	policy.Data = switcEntity
	// name policy
	if switcEntity.Name == "" && op == "insert" {
		policy.NameMessage = "Name is required"
	} else {
		// checking if name already exists
		if switcEntity.Name != "" {
			var dummyArray []entity.Switch
			filterEntity := entity.Switch{
				SwitchID: strings.Join(strings.Split(strings.ToLower(switcEntity.Name), " "), "-"),
			}
			dummyArray, _ = service.GetSwitch(ctx, filterEntity, "")
			if len(dummyArray) != 0 {
				policy.NameMessage = "This name is already exists"
			}
		}
	}

	// manufacturer policy
	if switcEntity.Manufacturer == "" && op == "insert" {
		policy.ManufacturerMessage = "Manufacturer is required"
	}
	// type policy
	if switcEntity.Type == "" && op == "insert" {
		policy.TypeMessage = "Type is required"
	} else {
		// checking if type correct
		if switcEntity.Type != "" {
			policy.TypeMessage = "Type must be either linear, clicky or silent"
			for _, tipe := range []string{"linear", "clicky", "silent"} {
				if switcEntity.Type == tipe {
					policy.TypeMessage = ""
					break
				}
			}
		}
	}

	if policy.NameMessage != "" || policy.ManufacturerMessage != "" || policy.TypeMessage != "" {
		return &policy
	}
	return nil
}
