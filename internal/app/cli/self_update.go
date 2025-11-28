package cli

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/geomark27/loom-go/internal/version"
	"github.com/spf13/cobra"
)

const (
	repoOwner  = "geomark27"
	repoName   = "loom-go"
	githubAPI  = "https://api.github.com/repos/%s/%s/releases/latest"
	githubTags = "https://api.github.com/repos/%s/%s/tags"
)

type githubRelease struct {
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
	HTMLURL string `json:"html_url"`
}

type githubTag struct {
	Name string `json:"name"`
}

var selfUpdateCmd = &cobra.Command{
	Use:   "self-update [version]",
	Short: "Update Loom CLI to the latest version",
	Long: `Update Loom CLI to the latest version or a specific version.

Examples:
  loom self-update           # Update to latest version
  loom self-update v1.1.2    # Update to specific version
  loom self-update --check   # Check for updates without installing`,
	Args: cobra.MaximumNArgs(1),
	RunE: runSelfUpdate,
}

var checkOnly bool

func init() {
	rootCmd.AddCommand(selfUpdateCmd)
	selfUpdateCmd.Flags().BoolVar(&checkOnly, "check", false, "Check for updates without installing")
}

func runSelfUpdate(cmd *cobra.Command, args []string) error {
	currentVersion := version.Current.String()
	fmt.Printf("🔧 Current version: v%s\n", currentVersion)

	var targetVersion string

	if len(args) > 0 {
		// Specific version requested
		targetVersion = args[0]
		if !strings.HasPrefix(targetVersion, "v") {
			targetVersion = "v" + targetVersion
		}
		fmt.Printf("📌 Target version: %s\n", targetVersion)
	} else {
		// Get latest version from GitHub
		fmt.Println("🔍 Checking for latest version...")
		latest, err := getLatestVersion()
		if err != nil {
			return fmt.Errorf("❌ Failed to check latest version: %w", err)
		}
		targetVersion = latest.TagName
		fmt.Printf("📦 Latest version: %s\n", targetVersion)
	}

	// Compare versions
	targetClean := strings.TrimPrefix(targetVersion, "v")
	if targetClean == currentVersion {
		fmt.Println("\n✅ You already have the latest version!")
		return nil
	}

	// Check if target is newer
	targetParsed, err := version.Parse(targetClean)
	if err != nil {
		return fmt.Errorf("❌ Invalid version format: %s", targetVersion)
	}

	if targetParsed.IsOlder(version.Current) {
		fmt.Printf("\n⚠️  Target version %s is older than current v%s\n", targetVersion, currentVersion)
		fmt.Println("   Use --force if you want to downgrade (not implemented yet)")
	}

	if checkOnly {
		if targetParsed.IsNewer(version.Current) {
			fmt.Printf("\n🆕 New version available: %s\n", targetVersion)
			fmt.Println("   Run 'loom self-update' to install")
		}
		return nil
	}

	// Perform update
	fmt.Printf("\n🚀 Updating to %s...\n\n", targetVersion)

	installCmd := exec.Command("go", "install",
		fmt.Sprintf("github.com/%s/%s/cmd/loom@%s", repoOwner, repoName, targetVersion))
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	installCmd.Env = append(os.Environ(),
		"GOPROXY=direct",
		fmt.Sprintf("GONOSUMDB=github.com/%s/%s", repoOwner, repoName),
	)

	if err := installCmd.Run(); err != nil {
		return fmt.Errorf("❌ Update failed: %w\n\nTry manually: go install github.com/%s/%s/cmd/loom@%s",
			repoOwner, repoName, targetVersion)
	}

	fmt.Printf("\n✅ Successfully updated to %s!\n", targetVersion)
	fmt.Println("   Run 'loom version' to verify")
	return nil
}

func getLatestVersion() (*githubRelease, error) {
	// Try releases first
	url := fmt.Sprintf(githubAPI, repoOwner, repoName)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var release githubRelease
		if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
			return nil, err
		}
		return &release, nil
	}

	// Fallback to tags if no releases
	return getLatestTag()
}

func getLatestTag() (*githubRelease, error) {
	url := fmt.Sprintf(githubTags, repoOwner, repoName)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
	}

	var tags []githubTag
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		return nil, err
	}

	if len(tags) == 0 {
		return nil, fmt.Errorf("no tags found")
	}

	// Find the latest semantic version tag
	var latestTag string
	var latestVersion version.Version

	for _, tag := range tags {
		if strings.HasPrefix(tag.Name, "v") {
			v, err := version.Parse(strings.TrimPrefix(tag.Name, "v"))
			if err == nil {
				if latestTag == "" || v.IsNewer(latestVersion) {
					latestTag = tag.Name
					latestVersion = v
				}
			}
		}
	}

	if latestTag == "" {
		return nil, fmt.Errorf("no valid version tags found")
	}

	return &githubRelease{TagName: latestTag}, nil
}
