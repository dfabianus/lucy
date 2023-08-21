/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package set

import (
	"fmt"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set called")
	},
}

func init() {
	//SetCmd.AddCommand(SetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	SetCmd.PersistentFlags().String("reactor-name", "rn", "define reactor name")
	SetCmd.PersistentFlags().String("reactor-id", "ri", "define reactor id")
	SetCmd.PersistentFlags().String("process-name", "pn", "define process name")
	SetCmd.PersistentFlags().String("process-id", "pi", "define processs id")
	SetCmd.PersistentFlags().String("port-name", "pon", "define port name")
	SetCmd.PersistentFlags().String("port-id", "poi", "define port id")
	SetCmd.PersistentFlags().String("value", "v", "define value to write")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
