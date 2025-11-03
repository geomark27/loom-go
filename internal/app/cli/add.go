package cli

import (
	"fmt"

	"github.com/geomark27/loom-go/internal/addon"
	"github.com/geomark27/loom-go/internal/generator"
	"github.com/spf13/cobra"
)

var (
	addForce bool
)

var addCmd = &cobra.Command{
	Use:   "add [type] [name]",
	Short: "Add components and technologies to the project",
	Long: `Add and configure new technologies in an existing Loom project.

Allows changing routers, adding ORMs, configuring databases,
implementing authentication, and more.

Available categories:
  router      - HTTP Frameworks (gin, chi, echo)
  orm         - ORMs (gorm, sqlc)
  database    - Databases (postgres, mysql, mongodb, redis)
  auth        - Authentication (jwt, oauth2)
  docker      - Containerization

Examples:
  loom add router gin          # Switch to Gin
  loom add orm gorm            # Add GORM
  loom add database postgres   # Configure PostgreSQL
  loom add auth jwt            # Add JWT auth
  loom add docker              # Add Dockerfile`,
	Args: cobra.MinimumNArgs(1),
	RunE: runAdd,
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().BoolVar(&addForce, "force", false, "Force installation (replaces existing)")
}

func runAdd(cmd *cobra.Command, args []string) error {
	if len(args) == 1 && args[0] == "list" {
		return showAvailableAddons()
	}

	if len(args) < 2 {
		return fmt.Errorf("usage: loom add [type] [name]\nExample: loom add router gin")
	}

	category := args[0]
	name := args[1]

	// Detect project
	projectInfo, err := generator.DetectProject()
	if err != nil {
		return fmt.Errorf("error: no valid Loom project detected. %w", err)
	}

	fmt.Printf("🔍 Project: %s (%s)\n", projectInfo.Name, projectInfo.Architecture)

	// Create addon manager
	manager := addon.NewAddonManager(projectInfo.RootPath, projectInfo.Architecture)

	// Map category to addon name
	addonName := mapCategoryToAddon(category, name)
	if addonName == "" {
		return fmt.Errorf("unrecognized addon: %s %s", category, name)
	}

	// Install addon
	fmt.Printf("📦 Adding %s %s...\n\n", category, name)

	if err := manager.InstallAddon(addonName, addForce); err != nil {
		return err
	}

	// Show next steps
	showNextSteps(category, name)

	return nil
}

func mapCategoryToAddon(category, name string) string {
	categories := map[string][]string{
		"router":   {"gin", "chi", "echo"},
		"orm":      {"gorm", "sqlc"},
		"database": {"postgres", "mysql", "mongodb", "redis"},
		"auth":     {"jwt", "oauth2"},
	}

	// Docker is special (no name)
	if category == "docker" {
		return "docker"
	}

	// Verify that the category exists
	validNames, exists := categories[category]
	if !exists {
		return ""
	}

	// Verify that the name is valid for that category
	for _, valid := range validNames {
		if name == valid {
			return name
		}
	}

	return ""
}

func showAvailableAddons() error {
	fmt.Println("📦 Available addons:")
	fmt.Println()

	fmt.Println("🌐 HTTP Routers:")
	fmt.Println("   loom add router gin      - Gin Web Framework")
	fmt.Println("   loom add router chi      - Chi Router")
	fmt.Println("   loom add router echo     - Echo Framework")

	fmt.Println("\n💾 ORMs:")
	fmt.Println("   loom add orm gorm        - GORM")
	fmt.Println("   loom add orm sqlc        - sqlc")

	fmt.Println("\n🗄️  Databases:")
	fmt.Println("   loom add database postgres   - PostgreSQL")
	fmt.Println("   loom add database mysql      - MySQL")
	fmt.Println("   loom add database mongodb    - MongoDB")
	fmt.Println("   loom add database redis      - Redis")

	fmt.Println("\n🔐 Authentication:")
	fmt.Println("   loom add auth jwt        - JWT Authentication")
	fmt.Println("   loom add auth oauth2     - OAuth 2.0")

	fmt.Println("\n🐳 Infrastructure:")
	fmt.Println("   loom add docker          - Docker + Docker Compose")

	fmt.Println("\n💡 Use 'loom add [type] [name]' to install")

	return nil
}

func showNextSteps(category, name string) {
	fmt.Println("\n📝 Next steps:")

	switch category {
	case "router":
		fmt.Println("   1. Run: go mod tidy")
		fmt.Println("   2. Update your handlers to use the new API")
		fmt.Println("   3. Run: go build ./cmd/...")

	case "orm":
		fmt.Println("   1. Run: go mod tidy")
		fmt.Println("   2. Configure the database connection")
		fmt.Println("   3. Update your repositories to use the ORM")

	case "database":
		fmt.Println("   1. Run: go mod tidy")
		fmt.Println("   2. Copy .env.example to .env and configure credentials")
		if name == "postgres" || name == "mysql" {
			fmt.Println("   3. Consider running: loom add docker")
		}

	case "auth":
		fmt.Println("   1. Run: go mod tidy")
		fmt.Println("   2. Copy .env.example to .env and change JWT_SECRET")
		fmt.Println("   3. Implement authentication endpoints")

	case "docker":
		fmt.Println("   1. Build the image: docker-compose build")
		fmt.Println("   2. Start containers: docker-compose up -d")
		fmt.Println("   3. View logs: docker-compose logs -f app")
	}

	fmt.Println("\n✨ Done! Your project has been updated")
}
