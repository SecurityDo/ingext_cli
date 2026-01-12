package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Manage streams",
}

var streamAddSourceCmd = &cobra.Command{
	Use:   "add-source",
	Short: "Add a stream source",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Adding stream datasource...")
		// Just call the global interface
		err := AppAPI.AddStreamSource()
		if err != nil {
			return err
		}

		fmt.Println("Stream source added successfully.")
		return nil
	},
}

// Example leaf command: source
var streamAddSinkCmd = &cobra.Command{
	Use:   "add-sink",
	Short: "Add a stream sink",
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
	streamCmd.AddCommand(streamAddSourceCmd, streamAddSinkCmd) // Add del/update similarly

	//streamAddCmd.AddCommand(streamAddSourceCmd)
	// Add other leaf commands: sink, router, connection
}
