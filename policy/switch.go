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
	policy.Data = switcEntity
	// name policy
	if switcEntity.Name == "" {
		policy.NameMessage = "Name is required"
	}
	// manufacturer policy
	if switcEntity.Manufacturer == "" {
		policy.ManufacturerMessage = "Manufacturer is required"
	}
	// type policy
	if switcEntity.Type == "" {
		policy.TypeMessage = "Type is required"
	} else {
		policy.TypeMessage = "Type must be either linear, clicky or silent"
		for _, tipe := range []string{"linear", "clicky", "silent"} {
			if switcEntity.Type == tipe {
				policy.TypeMessage = ""
				break
			}
		}
	}

	if policy.NameMessage != "" || policy.ManufacturerMessage != "" || policy.TypeMessage != "" {
		return &policy
	}
	return nil
}
