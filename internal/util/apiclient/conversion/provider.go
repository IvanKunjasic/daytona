// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package conversion

import (
	"github.com/daytonaio/daytona/internal/util"
	"github.com/daytonaio/daytona/pkg/apiclient"
	"github.com/daytonaio/daytona/pkg/models"
)

func ToApiClientProviderInfo(provider *models.ProviderInfo) *apiclient.ProviderInfo {
	targetConfigManifest := make(map[string]apiclient.TargetConfigProperty)

	for key, value := range provider.TargetConfigManifest {
		targetConfigManifest[key] = apiclient.TargetConfigProperty{
			DefaultValue:      &value.DefaultValue,
			Description:       &value.Description,
			DisabledPredicate: &value.DisabledPredicate,
			InputMasked:       &value.InputMasked,
			Options:           value.Options,
			Suggestions:       value.Suggestions,
			Type:              util.Pointer(apiclient.ModelsTargetConfigPropertyType(value.Type)),
		}
	}

	return &apiclient.ProviderInfo{
		AgentlessTarget:      &provider.AgentlessTarget,
		Label:                provider.Label,
		Name:                 provider.Name,
		RunnerId:             provider.RunnerId,
		Version:              provider.Version,
		TargetConfigManifest: targetConfigManifest,
	}
}

func ToProviderInfo(providerDto *apiclient.ProviderInfo) *models.ProviderInfo {
	targetConfigManifest := make(map[string]models.TargetConfigProperty)

	for key, value := range providerDto.TargetConfigManifest {
		targetConfigManifest[key] = models.TargetConfigProperty{
			DefaultValue:      *value.DefaultValue,
			Description:       *value.Description,
			DisabledPredicate: *value.DisabledPredicate,
			InputMasked:       *value.InputMasked,
			Options:           value.Options,
			Suggestions:       value.Suggestions,
			Type:              models.TargetConfigPropertyType(*value.Type),
		}
	}

	result := &models.ProviderInfo{
		Label:                providerDto.Label,
		Name:                 providerDto.Name,
		RunnerId:             providerDto.RunnerId,
		Version:              providerDto.Version,
		TargetConfigManifest: targetConfigManifest,
	}

	if providerDto.AgentlessTarget != nil {
		result.AgentlessTarget = *providerDto.AgentlessTarget
	}

	return result
}
