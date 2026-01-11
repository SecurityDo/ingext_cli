package commands

import (
	"fmt"
	"ingext/internal/config"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	confProvider string
	confContext  string
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the ingext tool",
	Long:  `Sets configuration values in ~/.ingext/config.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set values in Viper
		// Note: cluster and namespace are already bound in root.go,
		// but we need to explicitly set them if the user provided flags here to save them to file.

		if cluster != "" {
			viper.Set("cluster", cluster)
		}
		if namespace != "" {
			viper.Set("namespace", namespace)
		}
		if confProvider != "" {
			viper.Set("provider", confProvider)
		}
		if confContext != "" {
			viper.Set("context", confContext)
		}

		if err := config.SaveConfig(); err != nil {
			fmt.Printf("Error saving config: %v\n", err)
			return
		}

		fmt.Println("Configuration saved.")
	},
}

// 1. Define the 'view' subcommand
var configViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View current configuration settings",
	Long:  "Displays the current configuration loaded from ~/.ingext/config.yaml and environment variables.",
	Run: func(cmd *cobra.Command, args []string) {
		// Use tabwriter to create a clean, aligned table
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		fmt.Fprintln(w, "SETTING\tVALUE")
		fmt.Fprintln(w, "-------\t-----")

		// Retrieve values from Viper (checks flags, env vars, and config file)
		fmt.Fprintf(w, "Cluster\t%s\n", viper.GetString("cluster"))
		fmt.Fprintf(w, "Namespace\t%s\n", viper.GetString("namespace"))
		fmt.Fprintf(w, "Provider\t%s\n", viper.GetString("provider"))
		fmt.Fprintf(w, "Context\t%s\n", viper.GetString("context"))

		// Access config file location
		fmt.Fprintln(w, "-------\t-----")
		fmt.Fprintf(w, "Config File\t%s\n", viper.ConfigFileUsed())

		w.Flush()
	},
}

func init() {
	RootCmd.AddCommand(configCmd)

	// Local flags for the config command
	configCmd.Flags().StringVar(&confProvider, "provider", "", "Provider (eks|aks|gke)")
	configCmd.Flags().StringVar(&confContext, "context", "", "Kubeconfig context name")
	// Note: --cluster and --namespace are inherited from Root, but usually 'config' commands
	// might want to enforce them or treat them differently. For now, we rely on the Root persistent flags.

	// 2. Register 'view' as a child of 'config'
	configCmd.AddCommand(configViewCmd)
}
