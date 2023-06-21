// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/juju/jknife/ops"
)

var unitEnsureDeadCmd = &cobra.Command{
	Use:   "unit-ensure-dead",
	Short: "select a model",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if agentConfigPath == "" {
			return fmt.Errorf("--agent-config must be specified")
		}
		if modelUUID == "" {
			return fmt.Errorf("--model must be specified")
		}
		if unitName == "" {
			return fmt.Errorf("--unit must be specified")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return ops.UnitEnsureDead(agentConfigPath, modelUUID, unitName)
	},
}
