/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"github.com/spf13/cobra"
)

var ProcessFlag string
var SignalFlag string
var PortFlag string
var ReactorFlag string
var RunningFlag bool
var AllFlag bool

// getCmd represents the get command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Command to get information about reactors, processes and signals.",
	Long: `The get command is used to list information about reactors, processes and signals.
Combine the get command with one of the three subcommands and filter further by 
use of the provided flags. If the command and the filter flag is equal, the flag
can be ommited.

Examples:
lucy get reactors:			lists all available reactors in the setup
lucy get reactors --running: 		lists all reactors with running processes
lucy get reactors 20			get the reactor with id 20
lucy get reactors -p "test" 		get the reactor where process "test" was conducted
lucy get processes --running: 		lists all running processes
lucy get processes "test"		get the process with name "test"
lucy get processes "test" --all		list all processes with "test" in the name
lucy get processes 123			get the process with id 123
lucy get processes -s 115425		get the process of signal id 115425
lucy get signals -p "test" 		get the signals from process test`,
}

func init() {
	GetCmd.PersistentFlags().StringVarP(&ProcessFlag, "process", "p", "", "get by process name or id")
	GetCmd.PersistentFlags().StringVarP(&SignalFlag, "signal", "s", "", "get by signal id")
	GetCmd.PersistentFlags().StringVarP(&PortFlag, "port", "o", "", "get by port name or id")
	GetCmd.PersistentFlags().StringVarP(&ReactorFlag, "reactor", "r", "", "get by reactor name or id")
	GetCmd.PersistentFlags().BoolVarP(&RunningFlag, "running", "n", false, "filter to running status true")
	GetCmd.PersistentFlags().BoolVarP(&AllFlag, "all", "a", false, "fuzzy string search on")
}
