/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"lucy/lucrest"

	"github.com/spf13/cobra"
)

// reactorsCmd represents the reactors command

var reactorsCmd = &cobra.Command{
	Use:   "reactors",
	Short: "Reads reactors from the database",
	Long: `The get reactors function reads all available reactors from the database.
	It can be used to get reactor ids and running processes. 
	For a filter of running reactors use the -r, --running flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		if RunningFlag {
			lucrest.GetRunningReactors()
		} else {
			lucrest.GetReactors()
		}
	},
}

func init() {
	GetCmd.AddCommand(reactorsCmd)
}
