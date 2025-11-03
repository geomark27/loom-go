package version

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// DetectProjectVersion detects the Loom version used in the project
func DetectProjectVersion() (Version, error) {
	// Search for the Loom version comment in go.mod
	version, err := detectFromGoMod()
	if err == nil {
		return version, nil
	}

	// Search in .loom if it exists
	version, err = detectFromLoomFile()
	if err == nil {
		return version, nil
	}

	// If not found, assume oldest version
	return Version{Major: 1, Minor: 0, Patch: 0}, nil
}

// detectFromGoMod searches for the version in go.mod comments
func detectFromGoMod() (Version, error) {
	file, err := os.Open("go.mod")
	if err != nil {
		return Version{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	versionRegex := regexp.MustCompile(`(?i)loom\s+v?(\d+\.\d+\.\d+)`)

	for scanner.Scan() {
		line := scanner.Text()
		if matches := versionRegex.FindStringSubmatch(line); len(matches) > 1 {
			return Parse(matches[1])
		}
	}

	return Version{}, fmt.Errorf("version not found in go.mod")
}

// detectFromLoomFile reads the .loom file
func detectFromLoomFile() (Version, error) {
	content, err := os.ReadFile(".loom")
	if err != nil {
		return Version{}, err
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "version=") {
			versionStr := strings.TrimPrefix(line, "version=")
			return Parse(strings.TrimSpace(versionStr))
		}
	}

	return Version{}, fmt.Errorf("version not found in .loom")
}

// CreateLoomFile creates or updates the .loom file with the current version
func CreateLoomFile(version Version, architecture string) error {
	content := fmt.Sprintf(`# Loom Project Configuration
version=%s
architecture=%s
created_with=loom-cli
`, version.String(), architecture)

	return os.WriteFile(".loom", []byte(content), 0644)
}

// GetChangelogBetween returns the changelog between two versions
func GetChangelogBetween(from, to Version) string {
	// If both versions are 1.0.0 or higher, there are no internal changes
	if from.Major >= 1 && to.Major >= 1 {
		return "✅ Project updated. See CHANGELOG.md for complete details."
	}

	// For legacy projects (v0.x.x), suggest upgrade to v1.0.0
	if from.Major == 0 {
		return `🎉 Important update available!

📌 Version 1.0.0 - Stable Release:
  ✨ 'loom generate' command to create individual components
  🎨 'loom add' command to add technologies (routers, ORMs, databases)
  ⬆️ 'loom upgrade' command with versioning system
  📦 pkg/helpers updated and improved
  🏗️ Dual architecture (Layered + Modular)
  📚 Complete documentation renovated
  
� See CHANGELOG.md for complete details`
	}

	return "No changes recorded between these versions."
}
