package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "detect",
	Short: "A tool to get linux device information",
	Long:  `A tool to get linux device information like cpu, mem, and disk on different Linux distributions.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(cpuCmd)
	rootCmd.AddCommand(memCmd)
	rootCmd.AddCommand(allCmd)
	rootCmd.AddCommand(diskCmd)
}

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "show all information",
	Long:  `show all information like cpu, mem, and disk.`,
	Run: func(cmd *cobra.Command, args []string) {
		getCpuInfo()
		getMemInfo()
		getDiskInfo(rootRequired, homeRequired)
	},
}
