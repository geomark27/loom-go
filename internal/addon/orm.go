package addon

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ORMAddon manages ORM installation
type ORMAddon struct {
	projectRoot  string
	architecture string
	ormType      string // "gorm", "sqlc"
}

// NewORMAddon creates a new ORM addon
func NewORMAddon(projectRoot, architecture, ormType string) *ORMAddon {
	return &ORMAddon{
		projectRoot:  projectRoot,
		architecture: architecture,
		ormType:      ormType,
	}
}

func (o *ORMAddon) Name() string {
	return fmt.Sprintf("ORM %s", o.ormType)
}

func (o *ORMAddon) Description() string {
	descriptions := map[string]string{
		"gorm": "Complete ORM with advanced features",
		"sqlc": "Type-safe code generator from SQL",
	}
	return descriptions[o.ormType]
}

func (o *ORMAddon) IsInstalled() (bool, error) {
	detector := NewProjectDetector(o.projectRoot)
	currentORM := detector.DetectORM()
	return currentORM == o.ormType, nil
}

func (o *ORMAddon) CanInstall() (bool, string, error) {
	// Check that there's no other ORM installed
	detector := NewProjectDetector(o.projectRoot)
	currentORM := detector.DetectORM()

	if currentORM != "none" && currentORM != o.ormType {
		return false, fmt.Sprintf("You already have %s installed. Use --force to replace", currentORM), nil
	}

	return true, "", nil
}

func (o *ORMAddon) GetConflicts() []string {
	conflicts := []string{"gorm", "sqlc", "ent"}
	filtered := []string{}
	for _, c := range conflicts {
		if c != o.ormType {
			filtered = append(filtered, c)
		}
	}
	return filtered
}

func (o *ORMAddon) Install(force bool) error {
	switch o.ormType {
	case "gorm":
		return o.installGORM()
	case "sqlc":
		return o.installSQLC()
	default:
		return fmt.Errorf("unsupported ORM: %s", o.ormType)
	}
}

func (o *ORMAddon) installGORM() error {
	fmt.Println("   📦 Installing GORM...")

	// 1. Add GORM dependencies
	fmt.Println("   📦 Adding GORM dependencies...")
	deps := map[string]string{
		"gorm.io/gorm":            "v1.25.5",
		"gorm.io/driver/postgres": "v1.5.4",
		"golang.org/x/crypto":     "v0.17.0",
		"github.com/spf13/cobra":  "v1.9.1",
	}

	for module, version := range deps {
		if err := UpdateGoMod(module, version); err != nil {
			return fmt.Errorf("failed to add %s: %w", module, err)
		}
	}

	// 2. Create directory structure
	fmt.Println("   📁 Creating database structure...")
	dirs := []string{
		filepath.Join(o.projectRoot, "internal", "database"),
		filepath.Join(o.projectRoot, "internal", "database", "seeders"),
		filepath.Join(o.projectRoot, "cmd", "console"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// 3. Generate database files
	if err := o.generateDatabaseFiles(); err != nil {
		return err
	}

	// 4. Generate console command
	if err := o.generateConsoleCommand(); err != nil {
		return err
	}

	// 5. Update config to include database connection string
	if err := o.updateConfigForDatabase(); err != nil {
		return err
	}

	// 6. Update .env.example
	envVars := map[string]string{
		"DB_HOST":     "localhost",
		"DB_PORT":     "5432",
		"DB_NAME":     filepath.Base(o.projectRoot),
		"DB_USER":     "postgres",
		"DB_PASSWORD": "postgres",
		"DB_SSLMODE":  "disable",
	}

	if err := UpdateEnvExample(envVars, "Database Configuration"); err != nil {
		return err
	}

	// 7. Update Makefile
	if err := o.updateMakefileForDatabase(); err != nil {
		return err
	}

	fmt.Println("✅ GORM installed successfully!")
	fmt.Println("\n💡 Next steps:")
	fmt.Println("   1. Add database fields to your Config struct:")
	fmt.Println("      DBHost, DBPort, DBName, DBUser, DBPassword, DBSSLMode string")
	fmt.Println("   2. Add GetDBConnectionString() method to Config")
	fmt.Println("   3. Load DB environment variables in config.Load()")
	fmt.Println("   4. Update .env with database credentials")
	fmt.Println("   5. Run: go mod tidy")
	fmt.Println("   6. Run migrations: go run cmd/console/main.go migrate --seed")
	fmt.Println("   7. Or use Makefile: make db-migrate")
	fmt.Println("\n📝 See generated files in internal/database/ and cmd/console/")

	return nil
}

func (o *ORMAddon) installSQLC() error {
	fmt.Println("   📦 Installing sqlc...")

	// TODO: Implement sqlc installation
	fmt.Println("   ⚠️  Full implementation coming soon")
	fmt.Println("   💡 Visit https://docs.sqlc.dev/en/latest/overview/install.html")

	return nil
}

// generateDatabaseFiles creates database connection and model/seeder registry files
func (o *ORMAddon) generateDatabaseFiles() error {
	fmt.Println("   📝 Generating database files...")

	// Get module name from go.mod
	moduleName, err := GetModuleName(o.projectRoot)
	if err != nil {
		return fmt.Errorf("failed to get module name: %w", err)
	}

	// Determine config path based on architecture
	configPath := "internal/config"
	modelsPath := "internal/models"
	if o.architecture == "layered" {
		configPath = "internal/platform/config"
		modelsPath = "internal/app/models"
	}

	templates := map[string]string{
		"database.go":        "database/database.go.tmpl",
		"models_all.go":      "database/models_all.go.tmpl",
		"seeders_all.go":     "database/seeders_all.go.tmpl",
		"database_seeder.go": "database/database_seeder.go.tmpl",
		"user_seeder.go":     "database/user_seeder.go.tmpl",
	}

	for filename, tmplName := range templates {
		var targetPath string
		if filename == "user_seeder.go" {
			targetPath = filepath.Join(o.projectRoot, "internal", "database", "seeders", filename)
		} else if filename == "seeders_all.go" || filename == "database_seeder.go" {
			targetPath = filepath.Join(o.projectRoot, "internal", "database", "seeders", filename)
		} else {
			targetPath = filepath.Join(o.projectRoot, "internal", "database", filename)
		}

		if err := GenerateFileFromTemplate(tmplName, targetPath, map[string]interface{}{
			"ModuleName": moduleName,
			"ConfigPath": configPath,
			"ModelsPath": modelsPath,
		}); err != nil {
			return fmt.Errorf("failed to generate %s: %w", filename, err)
		}
	}

	return nil
}

// generateConsoleCommand creates the console CLI for migrations and seeders
func (o *ORMAddon) generateConsoleCommand() error {
	fmt.Println("   📝 Generating console command...")

	// Get module name from go.mod
	moduleName, err := GetModuleName(o.projectRoot)
	if err != nil {
		return fmt.Errorf("failed to get module name: %w", err)
	}

	// Determine paths based on architecture
	configPath := "internal/config"
	modelsPath := "internal/models"
	if o.architecture == "layered" {
		configPath = "internal/platform/config"
		modelsPath = "internal/app/models"
	}

	targetPath := filepath.Join(o.projectRoot, "cmd", "console", "main.go")
	return GenerateFileFromTemplate("console/main.go.tmpl", targetPath, map[string]interface{}{
		"ModuleName": moduleName,
		"ConfigPath": configPath,
		"ModelsPath": modelsPath,
	})
}

// updateConfigForDatabase adds database connection string method to config.go
func (o *ORMAddon) updateConfigForDatabase() error {
	fmt.Println("   ⚙️  Updating config for database...")
	fmt.Println("   ℹ️  Manual config update required (see instructions below)")

	// For now, we'll skip automatic config modification to avoid fragile string replacements
	// User will need to add DB fields manually or we can improve this in v1.2.0

	return nil
}

// updateMakefileForDatabase adds database-related targets to Makefile
func (o *ORMAddon) updateMakefileForDatabase() error {
	fmt.Println("   ⚙️  Updating Makefile...")

	makefilePath := filepath.Join(o.projectRoot, "Makefile")

	// Check if Makefile exists
	if _, err := os.Stat(makefilePath); os.IsNotExist(err) {
		fmt.Println("   ℹ️  No Makefile found, skipping")
		return nil
	}

	// Read current Makefile
	content, err := os.ReadFile(makefilePath)
	if err != nil {
		return fmt.Errorf("failed to read Makefile: %w", err)
	}

	makefileStr := string(content)

	// Check if database targets already exist
	if strings.Contains(makefileStr, "db-migrate") {
		fmt.Println("   ℹ️  Database targets already exist in Makefile")
		return nil
	}

	// Add database targets
	dbTargets := `
# Database commands
.PHONY: db-migrate db-seed db-fresh

db-migrate:
	@echo "Running migrations..."
	@go run cmd/console/main.go migrate

db-seed:
	@echo "Running seeders..."
	@go run cmd/console/main.go seed

db-fresh:
	@echo "Fresh migration with seeds..."
	@go run cmd/console/main.go migrate --fresh --seed
`

	// Append to the end of Makefile
	makefileStr += dbTargets

	// Write updated Makefile
	if err := os.WriteFile(makefilePath, []byte(makefileStr), 0644); err != nil {
		return fmt.Errorf("failed to write Makefile: %w", err)
	}

	return nil
}
