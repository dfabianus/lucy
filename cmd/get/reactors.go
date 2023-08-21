/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"lucy/lucrest"

	"github.com/spf13/cobra"
)

// reactorsCmd represents the reactors command
var RunningFlag bool

var reactorsCmd = &cobra.Command{
	Use:   "reactors",
	Short: "Reads reactors from the database",
	Long: `The get reactors function reads all available reactors from the database.
	It can be used to get reactor ids and running processes. 
	For a filter of running reactors use the -r, --running flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("reactors called")
		if RunningFlag {
			lucrest.GetRunningReactors()
		} else {
			lucrest.GetReactors()
		}
	},
}

func init() {
	//rootCmd.AddCommand(reactorsCmd)
	GetCmd.AddCommand(reactorsCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	reactorsCmd.PersistentFlags().BoolVarP(&RunningFlag, "running", "r", false, "filter to running status true")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reactorsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
