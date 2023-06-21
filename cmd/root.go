// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	agentConfigPath string
	modelUUID       string
	unitName        string
)

func init() {
	rootCmd.Flags().StringVarP(&agentConfigPath, "agent-config", "c", "", "Path to agent config found in /var/lib/juju")

	unitEnsureDeadCmd.Flags().AddFlagSet(rootCmd.Flags())
	unitEnsureDeadCmd.Flags().StringVarP(&modelUUID, "model", "m", "", "uuid of model to operate on")
	unitEnsureDeadCmd.Flags().StringVarP(&unitName, "unit", "u", "", "name of unit to operate on")
	rootCmd.AddCommand(unitEnsureDeadCmd)
}

var rootCmd = &cobra.Command{
	Use:   "jknife",
	Short: "jknife is a Juju surgical knife",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
