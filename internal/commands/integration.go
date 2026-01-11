package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	integType string
	integName string
)

var integrationCmd = &cobra.Command{
	Use:   "integration",
	Short: "Manage integrations",
}

var integrationAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an integration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Adding integration %s of type %s\n", integName, integType)
	},
}

var integrationDelCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete an integration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Deleting integration %s\n", integName)
	},
}

func init() {
	RootCmd.AddCommand(integrationCmd)
	integrationCmd.AddCommand(integrationAddCmd, integrationDelCmd)

	// Flags
	integrationAddCmd.Flags().StringVar(&integType, "integration", "", "Integration type")
	integrationAddCmd.Flags().StringVar(&integName, "name", "", "Name")
	_ = integrationAddCmd.MarkFlagRequired("integration")
	_ = integrationAddCmd.MarkFlagRequired("name")
}
