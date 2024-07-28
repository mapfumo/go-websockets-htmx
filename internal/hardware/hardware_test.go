package hardware

import (
	"strings"
	"testing"
)

// TestGetSystemSection checks if the GetSystemSection function
// returns a string containing all expected system information fields.
func TestGetSystemSection(t *testing.T) {
	output, err := GetSystemSection()
	if err != nil {
		t.Fatalf("GetSystemSection() error = %v", err)
	}

	// Define the fields we expect to see in the output
	expectedFields := []string{"Hostname:", "Total Memory:", "Used Memory:", "OS:"}

	// Check if each expected field is present in the output
	for _, field := range expectedFields {
		if !strings.Contains(output, field) {
			t.Errorf("GetSystemSection() output does not contain %s", field)
		}
	}
}

// TestGetCpuSection verifies that the GetCpuSection function
// returns a string containing the expected CPU information fields.
func TestGetCpuSection(t *testing.T) {
	output, err := GetCpuSection()
	if err != nil {
		t.Fatalf("GetCpuSection() error = %v", err)
	}

	// Define the fields we expect to see in the output
	expectedFields := []string{"CPU:", "Cores:"}

	// Check if each expected field is present in the output
	for _, field := range expectedFields {
		if !strings.Contains(output, field) {
			t.Errorf("GetCpuSection() output does not contain %s", field)
		}
	}
}

// TestGetDiskSection ensures that the GetDiskSection function
// returns a string containing the expected disk space information fields.
func TestGetDiskSection(t *testing.T) {
	output, err := GetDiskSection()
	if err != nil {
		t.Fatalf("GetDiskSection() error = %v", err)
	}

	// Define the fields we expect to see in the output
	expectedFields := []string{"Total Disk Space:", "Free Disk Space:"}

	// Check if each expected field is present in the output
	for _, field := range expectedFields {
		if !strings.Contains(output, field) {
			t.Errorf("GetDiskSection() output does not contain %s", field)
		}
	}
}

// TestAllSectionsIntegration is an integration test that checks if all
// hardware information functions work together correctly.
// It combines the output of all sections and verifies that all
// expected fields are present in the combined output.
func TestAllSectionsIntegration(t *testing.T) {
	// Get information from all sections
	systemInfo, err := GetSystemSection()
	if err != nil {
		t.Fatalf("GetSystemSection() error = %v", err)
	}

	cpuInfo, err := GetCpuSection()
	if err != nil {
		t.Fatalf("GetCpuSection() error = %v", err)
	}

	diskInfo, err := GetDiskSection()
	if err != nil {
		t.Fatalf("GetDiskSection() error = %v", err)
	}

	// Combine all information into a single string
	allInfo := strings.Join([]string{systemInfo, cpuInfo, diskInfo}, "\n\n")

	// Define all fields we expect to see in the combined output
	expectedFields := []string{
		"Hostname:", "Total Memory:", "Used Memory:", "OS:",
		"CPU:", "Cores:",
		"Total Disk Space:", "Free Disk Space:",
	}

	// Check if each expected field is present in the combined output
	for _, field := range expectedFields {
		if !strings.Contains(allInfo, field) {
			t.Errorf("Combined output does not contain %s", field)
		}
	}
}
