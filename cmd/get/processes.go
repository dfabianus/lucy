/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"lucy/lucrest"

	"github.com/spf13/cobra"
)

var ProcessNameFlag string
var ProcessIdFlag uint64

// processesCmd represents the processes command
var processesCmd = &cobra.Command{
	Use:   "processes",
	Short: "Reads processes from the database",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		lucrest.GetProcessByName(ProcessNameFlag)
		//lucrest.GetProcessById(ProcessIdFlag) doesn't work yet
	},
}

func init() {
	GetCmd.AddCommand(processesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// processesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	processesCmd.Flags().StringVarP(&ProcessNameFlag, "name", "n", "", "get processes by name")
	processesCmd.Flags().Uint64VarP(&ProcessIdFlag, "id", "i", 0, "get processes by process id")
}
