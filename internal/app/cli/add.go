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
	Use:   "add [tipo] [nombre]",
	Short: "Añade componentes y tecnologías al proyecto",
	Long: `Añade y configura nuevas tecnologías en un proyecto Loom existente.

Permite cambiar routers, añadir ORMs, configurar bases de datos,
implementar autenticación, y más.

Categorías disponibles:
  router      - Frameworks HTTP (gin, chi, echo)
  orm         - ORMs (gorm, sqlc)
  database    - Bases de datos (postgres, mysql, mongodb, redis)
  auth        - Autenticación (jwt, oauth2)
  docker      - Containerización

Ejemplos:
  loom add router gin          # Cambiar a Gin
  loom add orm gorm            # Agregar GORM
  loom add database postgres   # Configurar PostgreSQL
  loom add auth jwt            # Agregar JWT auth
  loom add docker              # Agregar Dockerfile`,
	Args: cobra.MinimumNArgs(1),
	RunE: runAdd,
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().BoolVar(&addForce, "force", false, "Forzar instalación (reemplaza existente)")
}

func runAdd(cmd *cobra.Command, args []string) error {
	if len(args) == 1 && args[0] == "list" {
		return showAvailableAddons()
	}

	if len(args) < 2 {
		return fmt.Errorf("uso: loom add [tipo] [nombre]\nEjemplo: loom add router gin")
	}

	category := args[0]
	name := args[1]

	// Detectar proyecto
	projectInfo, err := generator.DetectProject()
	if err != nil {
		return fmt.Errorf("error: no se detectó un proyecto Loom válido. %w", err)
	}

	fmt.Printf("🔍 Proyecto: %s (%s)\n", projectInfo.Name, projectInfo.Architecture)

	// Crear gestor de addons
	manager := addon.NewAddonManager(projectInfo.RootPath, projectInfo.Architecture)

	// Mapear categoría a nombre de addon
	addonName := mapCategoryToAddon(category, name)
	if addonName == "" {
		return fmt.Errorf("addon no reconocido: %s %s", category, name)
	}

	// Instalar addon
	fmt.Printf("📦 Añadiendo %s %s...\n\n", category, name)

	if err := manager.InstallAddon(addonName, addForce); err != nil {
		return err
	}

	// Mostrar próximos pasos
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

	// Docker es especial (no tiene nombre)
	if category == "docker" {
		return "docker"
	}

	// Verificar que la categoría existe
	validNames, exists := categories[category]
	if !exists {
		return ""
	}

	// Verificar que el nombre es válido para esa categoría
	for _, valid := range validNames {
		if name == valid {
			return name
		}
	}

	return ""
}

func showAvailableAddons() error {
	fmt.Println("📦 Addons disponibles:")
	fmt.Println()

	fmt.Println("🌐 Routers HTTP:")
	fmt.Println("   loom add router gin      - Gin Web Framework")
	fmt.Println("   loom add router chi      - Chi Router")
	fmt.Println("   loom add router echo     - Echo Framework")

	fmt.Println("\n💾 ORMs:")
	fmt.Println("   loom add orm gorm        - GORM")
	fmt.Println("   loom add orm sqlc        - sqlc")

	fmt.Println("\n🗄️  Bases de Datos:")
	fmt.Println("   loom add database postgres   - PostgreSQL")
	fmt.Println("   loom add database mysql      - MySQL")
	fmt.Println("   loom add database mongodb    - MongoDB")
	fmt.Println("   loom add database redis      - Redis")

	fmt.Println("\n🔐 Autenticación:")
	fmt.Println("   loom add auth jwt        - JWT Authentication")
	fmt.Println("   loom add auth oauth2     - OAuth 2.0")

	fmt.Println("\n🐳 Infrastructure:")
	fmt.Println("   loom add docker          - Docker + Docker Compose")

	fmt.Println("\n💡 Usa 'loom add [tipo] [nombre]' para instalar")

	return nil
}

func showNextSteps(category, name string) {
	fmt.Println("\n📝 Próximos pasos:")

	switch category {
	case "router":
		fmt.Println("   1. Ejecuta: go mod tidy")
		fmt.Println("   2. Actualiza tus handlers para usar la nueva API")
		fmt.Println("   3. Ejecuta: go build ./cmd/...")

	case "orm":
		fmt.Println("   1. Ejecuta: go mod tidy")
		fmt.Println("   2. Configura la conexión a la base de datos")
		fmt.Println("   3. Actualiza tus repositories para usar el ORM")

	case "database":
		fmt.Println("   1. Ejecuta: go mod tidy")
		fmt.Println("   2. Copia .env.example a .env y configura las credenciales")
		if name == "postgres" || name == "mysql" {
			fmt.Println("   3. Considera ejecutar: loom add docker")
		}

	case "auth":
		fmt.Println("   1. Ejecuta: go mod tidy")
		fmt.Println("   2. Copia .env.example a .env y cambia JWT_SECRET")
		fmt.Println("   3. Implementa los endpoints de autenticación")

	case "docker":
		fmt.Println("   1. Construye la imagen: docker-compose build")
		fmt.Println("   2. Inicia los containers: docker-compose up -d")
		fmt.Println("   3. Ve los logs: docker-compose logs -f app")
	}

	fmt.Println("\n✨ ¡Listo! Tu proyecto ha sido actualizado")
}
