package dcu

import (
	spec "github.com/compose-spec/compose-go/types"
)

type ComposeProject struct {
	*spec.Project
}

func (p *ComposeProject) ServiceMap() map[string]spec.ServiceConfig {
	serviceMap := make(map[string]spec.ServiceConfig)
	for _, service := range p.Services {
		serviceMap[service.Name] = service
	}
	return serviceMap
}
