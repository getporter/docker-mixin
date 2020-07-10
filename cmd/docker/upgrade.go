package main

import (
	"get.porter.sh/mixin/docker/pkg/docker"
	"github.com/spf13/cobra"
)

func buildUpgradeCommand(m *docker.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Execute the invoke functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Execute()
		},
	}
	return cmd
}
