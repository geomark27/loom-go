package addon

import "fmt"

// DatabaseAddon manages database configuration
type DatabaseAddon struct {
	projectRoot  string
	architecture string
	dbType       string // "postgres", "mysql", "mongodb", "redis"
}

// NewDatabaseAddon creates a new database addon
func NewDatabaseAddon(projectRoot, architecture, dbType string) *DatabaseAddon {
	return &DatabaseAddon{
		projectRoot:  projectRoot,
		architecture: architecture,
		dbType:       dbType,
	}
}

func (d *DatabaseAddon) Name() string {
	return fmt.Sprintf("Database %s", d.dbType)
}

func (d *DatabaseAddon) Description() string {
	descriptions := map[string]string{
		"postgres": "PostgreSQL - Robust relational database",
		"mysql":    "MySQL - Popular relational database",
		"mongodb":  "MongoDB - Document-oriented NoSQL database",
		"redis":    "Redis - In-memory database for caching",
	}
	return descriptions[d.dbType]
}

func (d *DatabaseAddon) IsInstalled() (bool, error) {
	detector := NewProjectDetector(d.projectRoot)
	databases := detector.DetectDatabase()

	for _, db := range databases {
		if db == d.dbType {
			return true, nil
		}
	}

	return false, nil
}

func (d *DatabaseAddon) CanInstall() (bool, string, error) {
	// Databases can coexist
	return true, "", nil
}

func (d *DatabaseAddon) GetConflicts() []string {
	return []string{} // Databases have no conflicts
}

func (d *DatabaseAddon) Install(force bool) error {
	switch d.dbType {
	case "postgres":
		return d.installPostgres()
	case "mysql":
		return d.installMySQL()
	case "mongodb":
		return d.installMongoDB()
	case "redis":
		return d.installRedis()
	default:
		return fmt.Errorf("unsupported database: %s", d.dbType)
	}
}

func (d *DatabaseAddon) installPostgres() error {
	fmt.Println("   📦 Configuring PostgreSQL...")

	// Add driver
	if err := UpdateGoMod("github.com/lib/pq", "v1.10.9"); err != nil {
		return err
	}

	// Update .env.example
	envVars := map[string]string{
		"DB_HOST":     "localhost",
		"DB_PORT":     "5432",
		"DB_USER":     "postgres",
		"DB_PASSWORD": "postgres",
		"DB_NAME":     "app_db",
		"DB_SSLMODE":  "disable",
	}

	if err := UpdateEnvExample(envVars, "PostgreSQL Database"); err != nil {
		return err
	}

	fmt.Println("   ✅ PostgreSQL configured")
	fmt.Println("   💡 Run 'loom add docker' to add PostgreSQL to docker-compose")

	return nil
}

func (d *DatabaseAddon) installMySQL() error {
	fmt.Println("   📦 Configuring MySQL...")

	// TODO: Implement MySQL
	fmt.Println("   ⚠️  Full implementation coming soon")

	return nil
}

func (d *DatabaseAddon) installMongoDB() error {
	fmt.Println("   📦 Configuring MongoDB...")

	// TODO: Implement MongoDB
	fmt.Println("   ⚠️  Full implementation coming soon")

	return nil
}

func (d *DatabaseAddon) installRedis() error {
	fmt.Println("   📦 Configuring Redis...")

	// TODO: Implement Redis
	fmt.Println("   ⚠️  Full implementation coming soon")

	return nil
}
