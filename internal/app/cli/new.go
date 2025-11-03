package cli

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/geomark27/loom-go/internal/generator"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new Go project with Loom",
	Long: `Create a new Go project following best practices and the
golang-standard/project-layout structure.

The generated project will include:
- Idiomatic directory structure
- Configured go.mod
- Basic web server with net/http
- README.md with instructions`,
	Args: cobra.ExactArgs(1),
	RunE: runNewCommand,
}

var (
	standalone bool
	moduleName string
	modular    bool
)

func runNewCommand(cmd *cobra.Command, args []string) error {
	projectName := args[0]

	// Validate project name
	if err := validateProjectName(projectName); err != nil {
		return fmt.Errorf("invalid project name: %w", err)
	}

	// Get current directory as base directory
	baseDir := "."
	projectPath := filepath.Join(baseDir, projectName)

	// Determine module name
	module := moduleName
	if module == "" {
		// Try to detect GitHub user from git config
		githubUser := detectGitHubUser()
		if githubUser != "" {
			module = fmt.Sprintf("github.com/%s/%s", githubUser, projectName)
		} else {
			// Fallback: use project name directly
			module = projectName
		}
	}

	// Determine architecture
	architecture := "layered"
	if modular {
		architecture = "modular"
	}

	// Create project configuration
	config := &generator.ProjectConfig{
		Name:         projectName,
		Path:         projectPath,
		ModuleName:   module,
		Description:  fmt.Sprintf("%s project generated with Loom", projectName),
		UseHelpers:   !standalone, // UseHelpers is true by default, false if --standalone is active
		IsModular:    modular,
		Architecture: architecture,
	}

	// Generate the project
	gen := generator.New()
	if err := gen.GenerateProject(config); err != nil {
		return fmt.Errorf("error generating project: %w", err)
	}

	// Success message with architecture information
	fmt.Printf("✅ Project '%s' created successfully in %s\n", projectName, projectPath)

	// Architecture information
	if config.IsModular {
		fmt.Printf("\n🏗️  Architecture: Modular (domain-based)\n")
		fmt.Printf("   → Ideal for: Large projects (20+ endpoints), teams, microservices\n")
		fmt.Printf("   → Modules: users (example generated)\n")
		fmt.Printf("\n💡 Tips:\n")
		fmt.Printf("   • Use 'loom generate module <name>' to add modules\n")
		fmt.Printf("   • Keep modules independent (use Event Bus for communication)\n")
		fmt.Printf("   • Each module has its own ports.go with interfaces\n")
	} else {
		fmt.Printf("\n🏗️  Architecture: Layered (layers-based)\n")
		fmt.Printf("   → Ideal for: Small APIs (< 20 endpoints), MVPs, prototypes\n")
		fmt.Printf("   → Structure: handlers → services → repositories\n")
		fmt.Printf("\n💡 Tips:\n")
		fmt.Printf("   • Start simple, scale when needed\n")
		fmt.Printf("   • Use 'loom generate module <name>' to add resources\n")
		fmt.Printf("   • Consider --modular if you have 20+ endpoints\n")
	}

	// Helpers information
	if config.UseHelpers {
		fmt.Printf("\n📦 Helpers: Included (validation, responses, logging)\n")
	} else {
		fmt.Printf("\n🔧 Mode: Standalone (no external dependencies)\n")
	}

	fmt.Printf("\nNext steps:\n")
	fmt.Printf("  cd %s\n", projectName)
	fmt.Printf("  go mod tidy\n")
	fmt.Printf("  go run cmd/%s/main.go\n", projectName)

	return nil
}

func validateProjectName(name string) error {
	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if strings.Contains(name, " ") {
		return fmt.Errorf("name cannot contain spaces")
	}

	// Check valid characters for directory names
	if strings.ContainsAny(name, `<>:"/\|?*`) {
		return fmt.Errorf("name contains invalid characters")
	}

	return nil
}

// detectGitHubUser tries to detect the GitHub username from git config
func detectGitHubUser() string {
	// Try to get github.user
	cmd := exec.Command("git", "config", "github.user")
	if output, err := cmd.Output(); err == nil {
		user := strings.TrimSpace(string(output))
		if user != "" {
			return user
		}
	}

	// Fallback: try to extract from remote origin URL
	cmd = exec.Command("git", "config", "remote.origin.url")
	if output, err := cmd.Output(); err == nil {
		url := strings.TrimSpace(string(output))
		// Parse URLs like: git@github.com:username/repo.git or https://github.com/username/repo.git
		if strings.Contains(url, "github.com") {
			// For SSH: git@github.com:username/repo.git
			if strings.HasPrefix(url, "git@github.com:") {
				parts := strings.Split(strings.TrimPrefix(url, "git@github.com:"), "/")
				if len(parts) > 0 {
					return parts[0]
				}
			}
			// For HTTPS: https://github.com/username/repo.git
			if strings.Contains(url, "github.com/") {
				parts := strings.Split(url, "github.com/")
				if len(parts) > 1 {
					userRepo := strings.Split(parts[1], "/")
					if len(userRepo) > 0 {
						return userRepo[0]
					}
				}
			}
		}
	}

	return ""
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Flags specific to the new command
	newCmd.Flags().StringVarP(&moduleName, "module", "m", "", "Go module name (auto-detects from git config or uses project name)")
	newCmd.Flags().BoolVar(&standalone, "standalone", false, "Generate project without Loom helpers (100% independent code)")
	newCmd.Flags().BoolVar(&modular, "modular", false, "Generate modular architecture by domain (recommended for large projects with 20+ endpoints)")
}
