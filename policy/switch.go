package policy

import "github.com/salamanderman234/peripheral-api/entity"

type SwitchPolicy struct {
	Data                entity.Switch `json:"data"`
	NameMessage         string        `json:"name,omitempty"`
	ManufacturerMessage string        `json:"manufacturer,omitempty"`
	TypeMessage         string        `json:"type,omitempty"`
}

func InsertSwitchPolicy(switcEntity entity.Switch) *SwitchPolicy {
	var policy SwitchPolicy
	var isReturn bool
	policy.Data = switcEntity
	if switcEntity.Name == "" {
		policy.NameMessage = "Name is required"
		isReturn = true
	}
	if switcEntity.Manufacturer == "" {
		policy.NameMessage = "Manufacturer is required"
		isReturn = true
	}
	if switcEntity.Type == "" {
		policy.NameMessage = "Type is required"
		isReturn = true
	}

	if isReturn {
		return &policy
	}
	return nil
}
