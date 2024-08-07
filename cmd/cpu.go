package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Get CPU information",
	Long:  `Get detailed information about the CPU, including model, number of cores, frequency, and current load.`,
	Run: func(cmd *cobra.Command, args []string) {
		getCpuInfo()
	},
}

func getCpuInfo() {
	fmt.Println("CPU Information:")
	// 获取CPU型号
	modelName := getModelName()
	fmt.Println("CPU Model:", modelName)

	// 获取CPU核心数
	cores := getCores()
	fmt.Println("Cores:", cores)

	// 获取CPU频率
	frequency := getFrequency()
	fmt.Println("Frequency (MHz):", frequency)

	// 获取当前核心负载
	load := getCpuLoad()
	fmt.Println("Current Load:", load)
	fmt.Println("")
	fmt.Println("")
}

func getModelName() string {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		fmt.Println("Error opening /proc/cpuinfo:", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "model name") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return "Unknown"
}

func getCores() int {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		fmt.Println("Error opening /proc/cpuinfo:", err)
		return 0
	}
	defer file.Close()

	cores := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "processor") {
			cores++
		}
	}
	return cores
}

func getFrequency() float64 {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		fmt.Println("Error opening /proc/cpuinfo:", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "cpu MHz") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				freq, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
				if err == nil {
					return freq
				}
			}
		}
	}
	return 0
}

func getCpuLoad() string {
	out, err := exec.Command("top", "-bn1").Output()
	if err != nil {
		fmt.Println("Error running top command:", err)
		return "Unknown"
	}
	re := regexp.MustCompile(`%Cpu\(s\):\s+([\d\.]+) us`)
	matches := re.FindStringSubmatch(string(out))
	if len(matches) > 1 {
		return matches[1] + " %"
	}
	return "Unknown"
}
