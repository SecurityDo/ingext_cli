package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Manage streams",
}

var streamAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a stream component",
}

// Example leaf command: source
var streamAddSourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Add a source",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Adding stream source...")
		// Just call the global interface
		err := AppAPI.AddStreamSource()
		if err != nil {
			return err
		}

		fmt.Println("Stream source added successfully.")
		return nil
	},
}

// ... Repeat for sink, router, connection ...

func init() {
	RootCmd.AddCommand(streamCmd)
	streamCmd.AddCommand(streamAddCmd) // Add del/update similarly

	streamAddCmd.AddCommand(streamAddSourceCmd)
	// Add other leaf commands: sink, router, connection
}
