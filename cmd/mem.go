package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var memCmd = &cobra.Command{
	Use:   "mem",
	Short: "Get memory information",
	Long:  `Get detailed information about the memory usage.`,
	Run: func(cmd *cobra.Command, args []string) {
		getMemInfo()
	},
}

func getMemInfo() {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		fmt.Println("Error opening /proc/meminfo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	memInfo := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) > 1 {
			memInfo[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}

	fmt.Println("Memory Information:")
	fmt.Println("Total Memory:", memInfo["MemTotal"])
	fmt.Println("Free Memory:", memInfo["MemFree"])
	fmt.Println("Available Memory:", memInfo["MemAvailable"])
	fmt.Println("Buffers:", memInfo["Buffers"])
	fmt.Println("Cached:", memInfo["Cached"])
	fmt.Println("")
	fmt.Println("")
}
