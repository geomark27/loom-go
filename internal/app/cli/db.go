package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/geomark27/loom-go/internal/generator"
	"github.com/spf13/cobra"
)

// Commands with colon syntax (loom db:migrate, loom db:fresh, loom db:seed)
var dbMigrateCmd = &cobra.Command{
	Use:   "db:migrate",
	Short: "Run database migrations",
	Long:  `Run all pending database migrations using GORM AutoMigrate.`,
	RunE:  runDBMigrate,
}

var dbFreshCmd = &cobra.Command{
	Use:   "db:fresh",
	Short: "Drop all tables and re-run migrations",
	Long: `Drop all tables registered in models_all.go and re-run migrations.
	
⚠️  WARNING: This is destructive! All data will be lost.
Use only in development environments.`,
	RunE: runDBFresh,
}

var dbSeedCmd = &cobra.Command{
	Use:   "db:seed",
	Short: "Run database seeders",
	Long:  `Execute all seeders registered in seeders_all.go.`,
	RunE:  runDBSeed,
}

var (
	seedAfterFresh   bool
	seedAfterMigrate bool
)

func init() {
	// Add commands directly to root (loom db:migrate, loom db:fresh, loom db:seed)
	rootCmd.AddCommand(dbMigrateCmd)
	rootCmd.AddCommand(dbFreshCmd)
	rootCmd.AddCommand(dbSeedCmd)

	// Add --seed flag to migrate and fresh
	dbMigrateCmd.Flags().BoolVar(&seedAfterMigrate, "seed", false, "Run seeders after migration")
	dbFreshCmd.Flags().BoolVar(&seedAfterFresh, "seed", false, "Run seeders after fresh migration")
}

func runDBMigrate(cmd *cobra.Command, args []string) error {
	return executeConsoleCommand("migrate", seedAfterMigrate, false)
}

func runDBFresh(cmd *cobra.Command, args []string) error {
	return executeConsoleCommand("migrate", seedAfterFresh, true)
}

func runDBSeed(cmd *cobra.Command, args []string) error {
	return executeConsoleCommand("seed", false, false)
}

func executeConsoleCommand(command string, withSeed bool, fresh bool) error {
	// Detect project
	projectInfo, err := generator.DetectProject()
	if err != nil {
		return fmt.Errorf("❌ Not in a Loom project: %w", err)
	}

	// Check if console exists
	consolePath := filepath.Join(projectInfo.RootPath, "cmd", "console", "main.go")
	if _, err := os.Stat(consolePath); os.IsNotExist(err) {
		return fmt.Errorf(`❌ Console CLI not found.

Run 'loom add orm gorm' first to generate the database structure.`)
	}

	// Build command arguments
	cmdArgs := []string{"run", "cmd/console/main.go", command}

	if fresh {
		cmdArgs = append(cmdArgs, "--fresh")
	}

	if withSeed {
		cmdArgs = append(cmdArgs, "--seed")
	}

	// Execute go run cmd/console/main.go <command>
	fmt.Printf("🚀 Running: go %s\n\n", joinArgs(cmdArgs))

	execCmd := exec.Command("go", cmdArgs...)
	execCmd.Dir = projectInfo.RootPath
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	execCmd.Stdin = os.Stdin

	return execCmd.Run()
}

func joinArgs(args []string) string {
	result := ""
	for i, arg := range args {
		if i > 0 {
			result += " "
		}
		result += arg
	}
	return result
}
