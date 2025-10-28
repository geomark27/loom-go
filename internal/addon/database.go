package addon

import "fmt"

// DatabaseAddon gestiona la configuración de bases de datos
type DatabaseAddon struct {
	projectRoot  string
	architecture string
	dbType       string // "postgres", "mysql", "mongodb", "redis"
}

// NewDatabaseAddon crea un nuevo addon de base de datos
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
		"postgres": "PostgreSQL - Base de datos relacional robusta",
		"mysql":    "MySQL - Base de datos relacional popular",
		"mongodb":  "MongoDB - Base de datos NoSQL orientada a documentos",
		"redis":    "Redis - Base de datos en memoria para cache",
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
	// Las bases de datos pueden coexistir
	return true, "", nil
}

func (d *DatabaseAddon) GetConflicts() []string {
	return []string{} // Las bases de datos no tienen conflictos
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
		return fmt.Errorf("base de datos no soportada: %s", d.dbType)
	}
}

func (d *DatabaseAddon) installPostgres() error {
	fmt.Println("   📦 Configurando PostgreSQL...")

	// Añadir driver
	if err := UpdateGoMod("github.com/lib/pq", "v1.10.9"); err != nil {
		return err
	}

	// Actualizar .env.example
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

	fmt.Println("   ✅ PostgreSQL configurado")
	fmt.Println("   💡 Ejecuta 'loom add docker' para añadir PostgreSQL a docker-compose")

	return nil
}

func (d *DatabaseAddon) installMySQL() error {
	fmt.Println("   📦 Configurando MySQL...")

	// TODO: Implementar MySQL
	fmt.Println("   ⚠️  Implementación completa próximamente")

	return nil
}

func (d *DatabaseAddon) installMongoDB() error {
	fmt.Println("   📦 Configurando MongoDB...")

	// TODO: Implementar MongoDB
	fmt.Println("   ⚠️  Implementación completa próximamente")

	return nil
}

func (d *DatabaseAddon) installRedis() error {
	fmt.Println("   📦 Configurando Redis...")

	// TODO: Implementar Redis
	fmt.Println("   ⚠️  Implementación completa próximamente")

	return nil
}
