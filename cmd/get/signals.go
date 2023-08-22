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
		lucrest.GetSignalsByProcessId(ProcessIdFlag)
		// if SignalIdFlagSignals != 0 {
		// 	lucrest.GetSignalsBySignalId(SignalIdFlagSignals)
		// }
	},
}

func init() {
	GetCmd.AddCommand(signalsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signalsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signalsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// signalsCmd.Flags().Uint64VarP(&SignalIdFlagSignals, "signal-id", "s", 0, "get signals by process id")
}
