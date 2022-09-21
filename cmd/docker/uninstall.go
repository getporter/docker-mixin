package main

import (
	"get.porter.sh/mixin/docker/pkg/docker"
	"github.com/spf13/cobra"
)

func buildUninstallCommand(m *docker.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Execute the uninstall functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Execute(cmd.Context())
		},
	}
	return cmd
}
