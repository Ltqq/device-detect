package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"strconv"
	"strings"
)

var rootRequired int
var homeRequired int

func init() {
	diskCmd.Flags().IntVarP(&rootRequired, "root", "r", 0, "Required space for root directory (GB)")
	diskCmd.Flags().IntVarP(&homeRequired, "home", "m", 0, "Required space for home directory (GB)")

}

var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Get disk information",
	Long:  `Get detailed information about disk usage.`,
	Run: func(cmd *cobra.Command, args []string) {
		getDiskInfo(rootRequired, homeRequired)
	},
}

func getDiskInfo(rootRequired int, homeRequired int) {
	fmt.Println("DISK Information:")
	out, err := exec.Command("df", "-h").Output()
	if err != nil {
		fmt.Println("Error running df command:", err)
		return
	}

	rootAvailable, homeAvailable := parseDiskInfo(out)

	fmt.Printf("Root Directory Available Space: %d GB\n", rootAvailable)
	fmt.Printf("Home Directory Available Space: %d GB\n", homeAvailable)

	if rootRequired > 0 {
		fmt.Printf("Root Directory Requirement: %d GB - ", rootRequired)
		if rootAvailable >= rootRequired {
			fmt.Println("Satisfied")
		} else {
			fmt.Println("Not Satisfied")
		}
	}

	if homeRequired > 0 {
		fmt.Printf("Home Directory Requirement: %d GB - ", homeRequired)
		if homeAvailable >= homeRequired {
			fmt.Println("Satisfied")
		} else {
			fmt.Println("Not Satisfied")
		}
	}
	fmt.Println("")
	fmt.Println("")
}

func parseDiskInfo(output []byte) (int, int) {
	scanner := bufio.NewScanner(bytes.NewReader(output))
	var rootAvailable, homeAvailable int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "/dev/") {
			parts := strings.Fields(line)
			if len(parts) >= 6 {
				mountPoint := parts[5]
				availableSpace := parts[3]
				if strings.HasSuffix(availableSpace, "G") {
					space, err := strconv.Atoi(strings.TrimSuffix(availableSpace, "G"))
					if err == nil {
						if mountPoint == "/" {
							rootAvailable = space
						} else if mountPoint == "/home" {
							homeAvailable = space
						}
					}
				}
			}
		}
	}

	// 如果/home不可用，查找是否有同一个文件系统
	if homeAvailable == 0 {
		homeAvailable = rootAvailable
	}

	return rootAvailable, homeAvailable
}
