// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package target

import (
	"context"

	"github.com/daytonaio/daytona/cmd/daytona/config"
	"github.com/daytonaio/daytona/internal/util"
	apiclient_util "github.com/daytonaio/daytona/internal/util/apiclient"
	"github.com/daytonaio/daytona/pkg/apiclient"
	target_util "github.com/daytonaio/daytona/pkg/cmd/target/util"
	"github.com/daytonaio/daytona/pkg/ide"
	"github.com/daytonaio/daytona/pkg/views/target/selection"
	views_util "github.com/daytonaio/daytona/pkg/views/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var sshOptions []string

var SshCmd = &cobra.Command{
	Use:     "ssh [TARGET] [WORKSPACE] [CMD...]",
	Short:   "SSH into a workspace using the terminal",
	Args:    cobra.ArbitraryArgs,
	GroupID: util.TARGET_GROUP,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := config.GetConfig()
		if err != nil {
			return err
		}

		activeProfile, err := c.GetActiveProfile()
		if err != nil {
			return err
		}

		ctx := context.Background()
		var target *apiclient.TargetDTO
		var workspaceName string
		var providerConfigId *string

		apiClient, err := apiclient_util.GetApiClient(&activeProfile)
		if err != nil {
			return err
		}

		if len(args) == 0 {
			targetList, res, err := apiClient.TargetAPI.ListTargets(ctx).Execute()
			if err != nil {
				return apiclient_util.HandleErrorResponse(res, err)
			}

			if len(targetList) == 0 {
				views_util.NotifyEmptyTargetList(true)
				return nil
			}

			target = selection.GetTargetFromPrompt(targetList, "SSH Into")
			if target == nil {
				return nil
			}
		} else {
			target, err = apiclient_util.GetTarget(args[0], true)
			if err != nil {
				return err
			}
		}

		if len(args) == 0 || len(args) == 1 {
			selectedWorkspace, err := selectTargetWorkspace(target.Id, &activeProfile)
			if err != nil {
				return err
			}
			if selectedWorkspace == nil {
				return nil
			}
			workspaceName = selectedWorkspace.Name
			providerConfigId = selectedWorkspace.GitProviderConfigId
		}

		if len(args) >= 2 {
			workspaceName = args[1]
			for _, workspace := range target.Workspaces {
				if workspace.Name == workspaceName {
					providerConfigId = workspace.GitProviderConfigId
					break
				}
			}
		}

		if !target_util.IsWorkspaceRunning(target, workspaceName) {
			wsRunningStatus, err := AutoStartTarget(target.Name, workspaceName)
			if err != nil {
				return err
			}
			if !wsRunningStatus {
				return nil
			}
		}

		sshArgs := []string{}
		if len(args) > 2 {
			sshArgs = append(sshArgs, args[2:]...)
		}

		gpgKey, err := GetGitProviderGpgKey(apiClient, ctx, providerConfigId)
		if err != nil {
			log.Warn(err)
		}

		return ide.OpenTerminalSsh(activeProfile, target.Id, workspaceName, gpgKey, sshOptions, sshArgs...)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) >= 2 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		if len(args) == 1 {
			return getWorkspaceNameCompletions(cmd, args, toComplete)
		}

		return getTargetNameCompletions()
	},
}

func init() {
	SshCmd.Flags().BoolVarP(&yesFlag, "yes", "y", false, "Automatically confirm any prompts")
	SshCmd.Flags().StringArrayVarP(&sshOptions, "option", "o", []string{}, "Specify SSH options in KEY=VALUE format.")
}