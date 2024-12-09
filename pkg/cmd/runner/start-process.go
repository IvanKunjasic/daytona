// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package runner

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var startProcessRunnerCmd = &cobra.Command{
	Use:    "start-process",
	Short:  "Starts the runner in the foreground",
	Hidden: true,
	Args:   cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if log.GetLevel() < log.InfoLevel {
			log.SetLevel(log.InfoLevel)
		}

		log.Info("Starting the runner...")

		// err := bootstrap.InitProviderManager(params.ServerConfig, params.RunnerConfig, params.ConfigDir)
		// if err != nil {
		// 	return err
		// }

		// runner, err := bootstrap.GetRemoteRunner(params)
		// if err != nil {
		// 	return err
		// }
		return nil
	},
}
