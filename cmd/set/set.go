/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package set

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ProcessFlag string
var SignalFlag string
var PortFlag string
var ReactorFlag string
var RunningFlag bool

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
	SetCmd.PersistentFlags().StringVarP(&ProcessFlag, "process", "p", "", "set by process name or id")
	SetCmd.PersistentFlags().StringVarP(&SignalFlag, "signal", "s", "", "set by signal id")
	SetCmd.PersistentFlags().StringVarP(&PortFlag, "port", "o", "", "set by port name or id")
	SetCmd.PersistentFlags().StringVarP(&ReactorFlag, "reactor", "r", "", "set by reactor name or id")
}
