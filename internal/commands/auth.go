package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	authName        string
	authDisplayName string
	authRole        string
)

// Parent command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage authentication (users and tokens)",
}

// Verbs
var authAddCmd = &cobra.Command{Use: "add", Short: "Add a resource"}
var authDelCmd = &cobra.Command{Use: "del", Short: "Delete a resource"}
var authUpdateCmd = &cobra.Command{Use: "update", Short: "Update a resource"}

// Nouns (User)
var authAddUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Add a new user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Adding user: %s (Role: %s)\n", authName, authRole)
	},
}

// Nouns (Token)
var authAddTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Add a new token",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Adding token for: %s\n", authName)
	},
}

// ... Repeat similar logic for del/update user/token ...

func init() {
	RootCmd.AddCommand(authCmd)
	authCmd.AddCommand(authAddCmd, authDelCmd, authUpdateCmd)

	// Add 'user' and 'token' to 'add'
	authAddCmd.AddCommand(authAddUserCmd, authAddTokenCmd)
	
	// Add flags to the leaf commands (or persistent flags to the verbs)
	authAddUserCmd.Flags().StringVar(&authName, "name", "", "Name of the user")
	authAddUserCmd.Flags().StringVar(&authDisplayName, "displayName", "", "Display name")
	authAddUserCmd.Flags().StringVar(&authRole, "role", "", "Role (admin|analyst)")
	
	// Mark required
	_ = authAddUserCmd.MarkFlagRequired("name")
}
