package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var registry, login, password string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "docker-auth",
	Short: "Tool to add docker registry credentials to docker config",
	Long:  `Tool to add docker registry credentials to docker config`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	addCmd.Flags().StringVar(&registry, "r", "", "registry")
	addCmd.Flags().StringVar(&login, "l", "", "login of the registry")
	addCmd.Flags().StringVar(&password, "p", "", "password of the registry")
	rootCmd.AddCommand(addCmd)
}
