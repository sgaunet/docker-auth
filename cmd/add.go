package cmd

import (
	"fmt"
	"os"

	"github.com/sgaunet/docker-auth/pkg/dockerauth"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add docker registry credentials to docker config",
	Long:  `add docker registry credentials to docker config`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		if login == "" {
			fmt.Fprintf(os.Stderr, "login not specified\n")
			_ = cmd.Help()
			os.Exit(1)
		}

		if password == "" {
			fmt.Fprintf(os.Stderr, "password not specified\n")
			_ = cmd.Help()
			os.Exit(1)
		}

		if registry == "" {
			fmt.Fprintf(os.Stderr, "registry not specified\n")
			_ = cmd.Help()
			os.Exit(1)
		}

		configFile := dockerauth.DefaultConfigFile
		payload, err := dockerauth.LoadDockerConfig(configFile)
		if err != nil {
			if err != os.ErrNotExist {
				fmt.Fprintf(os.Stderr, "error loading docker config: %v\n", err)
				os.Exit(1)
			}
		}
		err = dockerauth.AddAuthToDockerConfig(payload, registry, login, password)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error adding auth to docker config: %v\n", err)
			os.Exit(1)
		}
		err = dockerauth.SaveDockerConfig(configFile, payload)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error saving docker config: %v\n", err)
			os.Exit(1)
		}
	},
}
