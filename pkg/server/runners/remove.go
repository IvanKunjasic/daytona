// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package runners

import (
	"context"

	"github.com/daytonaio/daytona/pkg/stores"

	log "github.com/sirupsen/logrus"
)

func (s *RunnerService) RemoveRunner(ctx context.Context, runnerId string) error {
	var err error
	ctx, err = s.runnerStore.BeginTransaction(ctx)
	if err != nil {
		return s.runnerStore.RollbackTransaction(ctx, err)
	}

	defer stores.RecoverAndRollback(ctx, s.runnerStore)

	runner, err := s.runnerStore.Find(ctx, runnerId)
	if err != nil {
		return s.runnerStore.RollbackTransaction(ctx, err)
	}

	err = s.runnerStore.Delete(ctx, runner)
	if err != nil {
		return s.runnerStore.RollbackTransaction(ctx, err)
	}

	metadata, err := s.runnerMetadataStore.Find(ctx, runnerId)
	if err != nil && !stores.IsRunnerMetadataNotFound(err) {
		return s.runnerStore.RollbackTransaction(ctx, err)
	}
	if metadata != nil {
		err = s.runnerMetadataStore.Delete(ctx, metadata)
		if err != nil {
			return s.runnerStore.RollbackTransaction(ctx, err)
		}
	}

	err = s.revokeApiKey(ctx, runner.Alias)
	if err != nil {
		// Should not fail the whole operation if the API key cannot be revoked
		log.Error(err)
	}

	err = s.runnerStore.CommitTransaction(ctx)
	return s.runnerStore.RollbackTransaction(ctx, err)
}
