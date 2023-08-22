/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"lucy/lucrest"

	"github.com/spf13/cobra"
)

// signalsCmd represents the signals command
var signalsCmd = &cobra.Command{
	Use:   "signals",
	Short: "Reads signals from the database",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		lucrest.GetSignalsByProcessId(ProcessFlag)
	},
}

func init() {
	GetCmd.AddCommand(signalsCmd)
}
