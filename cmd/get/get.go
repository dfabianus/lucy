/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"github.com/spf13/cobra"
)

var ProcessIdFlag uint64
var ProcessNameFlag string
var SignalIdFlag uint64
var PortIdFlag uint64
var PortNameFlag string

// getCmd represents the get command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("get called")
	// },
}

func init() {
	//rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")
	GetCmd.PersistentFlags().Uint64VarP(&ProcessIdFlag, "process-id", "p", 0, "get by process id")
	GetCmd.PersistentFlags().StringVarP(&ProcessNameFlag, "process-name", "n", "", "get by process name")
	GetCmd.PersistentFlags().Uint64VarP(&SignalIdFlag, "signal-id", "s", 0, "get by signal id")
	GetCmd.PersistentFlags().Uint64VarP(&PortIdFlag, "port-id", "o", 0, "get by port id")
	GetCmd.PersistentFlags().StringVarP(&PortNameFlag, "port-name", "f", "", "get by port name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
