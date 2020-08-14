package main

import (
	"fmt"
	"io"
	"os"

	"get.porter.sh/mixin/docker/pkg/docker"
	"github.com/spf13/cobra"
)

func main() {
	cmd, err := buildRootCommand(os.Stdin)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		os.Exit(1)
	}
	if err := cmd.Execute(); err != nil {
		fmt.Printf("err: %s\n", err)
		os.Exit(1)
	}
}

func buildRootCommand(in io.Reader) (*cobra.Command, error) {
	m := docker.New()
	m.In = in
	cmd := &cobra.Command{
		Use:  "docker",
		Long: "Use the docker CLI with Porter",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Enable swapping out stdout/stderr for testing
			m.Out = cmd.OutOrStdout()
			m.Err = cmd.OutOrStderr()
		},
		SilenceUsage: true,
	}

	cmd.PersistentFlags().BoolVar(&m.Debug, "debug", false, "Enable debug logging")

	cmd.AddCommand(buildVersionCommand(m))
	cmd.AddCommand(buildSchemaCommand(m))
	cmd.AddCommand(buildBuildCommand(m))
	cmd.AddCommand(buildInstallCommand(m))
	cmd.AddCommand(buildInvokeCommand(m))
	cmd.AddCommand(buildUpgradeCommand(m))
	cmd.AddCommand(buildUninstallCommand(m))

	return cmd, nil
}
