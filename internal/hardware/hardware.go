package hardware

import (
	"fmt"
	"runtime"

	cpu "github.com/shirou/gopsutil/cpu"
	disk "github.com/shirou/gopsutil/disk"
	host "github.com/shirou/gopsutil/host"
	mem "github.com/shirou/gopsutil/mem"
)

// GetSystemSection gathers system information such as hostname, total and used memory, and OS.
// Returns the information as a formatted string and any error encountered.
func GetSystemSection() (string, error) {
	runTimeOS := runtime.GOOS          // Get the operating system
	vmStat, err := mem.VirtualMemory() // Get virtual memory statistics
	if err != nil {
		return "", err // Return error if there's any
	}

	hostStat, err := host.Info() // Get host information
	if err != nil {
		return "", err // Return error if there's any
	}

	// Format the output string with the collected information
	output := fmt.Sprintf("Hostname: %s<br>\nTotal Memory: %d<br>\nUsed Memory: %d<br>OS: %s<hr>",
		hostStat.Hostname,
		vmStat.Total,
		vmStat.Used,
		runTimeOS,
	)

	return output, nil
}

// GetCpuSection gathers CPU information such as model name and the number of cores.
// Returns the information as a formatted string and any error encountered.
func GetCpuSection() (string, error) {
	cpuStat, err := cpu.Info() // Get CPU information
	if err != nil {
		return "", err // Return error if there's any
	}
	// Format the output string with the collected information
	output := fmt.Sprintf("CPU: %s<br>\nCores: %d<hr>", cpuStat[0].ModelName, len(cpuStat))
	return output, nil
}

// GetDiskSection gathers disk usage information such as total and free disk space.
// Returns the information as a formatted string and any error encountered.
func GetDiskSection() (string, error) {
	diskStat, err := disk.Usage("/") // Get disk usage information
	if err != nil {
		return "", err // Return error if there's any
	}
	// Format the output string with the collected information
	output := fmt.Sprintf("Total Disk Space: %d<br>\nFree Disk Space: %d", diskStat.Total, diskStat.Free)
	return output, nil
}
