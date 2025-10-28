package cli

import (
	"fmt"

	"github.com/geomark27/loom-go/internal/generator"
	"github.com/spf13/cobra"
)

var generateModuleCmd = &cobra.Command{
	Use:   "module [nombre]",
	Short: "Genera un módulo completo (handler, service, repository, model, DTO)",
	Long: `Genera un módulo completo con todas sus capas en el proyecto actual.

El comando detectará automáticamente la arquitectura de tu proyecto
(Layered o Modular) y generará los archivos apropiados.

Para arquitectura Layered, genera:
  - internal/app/handlers/{nombre}_handler.go
  - internal/domain/services/{nombre}_service.go
  - internal/infrastructure/repositories/{nombre}_repository.go
  - internal/domain/models/{nombre}.go
  - internal/domain/dto/{nombre}_dto.go

Para arquitectura Modular, genera:
  - internal/modules/{nombre}/handler.go
  - internal/modules/{nombre}/service.go
  - internal/modules/{nombre}/repository.go
  - internal/modules/{nombre}/model.go
  - internal/modules/{nombre}/dto.go
  - internal/modules/{nombre}/router.go
  - internal/modules/{nombre}/validator.go
  - internal/modules/{nombre}/errors.go

Ejemplos:
  loom generate module products
  loom generate module users --force
  loom generate module orders --dry-run`,
	Aliases: []string{"mod", "m"},
	Args:    cobra.ExactArgs(1),
	RunE:    runGenerateModule,
}

func init() {
	generateCmd.AddCommand(generateModuleCmd)
}

func runGenerateModule(cmd *cobra.Command, args []string) error {
	moduleName := args[0]
	force, _ := cmd.Flags().GetBool("force")
	dryRun, _ := cmd.Flags().GetBool("dry-run")

	// Detectar el proyecto actual (sin argumentos)
	projectInfo, err := generator.DetectProject()
	if err != nil {
		return fmt.Errorf("error: no se detectó un proyecto Loom válido. %w", err)
	}

	// Validar el nombre del módulo
	if err := generator.ValidateComponentName(moduleName); err != nil {
		return fmt.Errorf("nombre de módulo inválido: %w", err)
	}

	fmt.Printf("🔍 Proyecto detectado: %s\n", projectInfo.Name)
	fmt.Printf("📐 Arquitectura: %s\n", projectInfo.Architecture)
	fmt.Printf("📦 Generando módulo: %s\n\n", moduleName)

	// Crear el generador
	gen := generator.NewModuleGenerator(projectInfo)

	// Generar el módulo (devuelve la lista de archivos)
	files, err := gen.GenerateModule(moduleName, force, dryRun)
	if err != nil {
		return fmt.Errorf("error al generar módulo: %w", err)
	}

	// Modo dry-run
	if dryRun {
		fmt.Println("📋 Archivos que se generarían:")
		for _, file := range files {
			fmt.Printf("   ✨ %s\n", file)
		}
		fmt.Println("\n💡 Ejecuta sin --dry-run para crear los archivos")
		return nil
	}

	fmt.Println("✅ Módulo generado exitosamente!")
	fmt.Println("\n📝 Archivos creados:")
	for _, file := range files {
		fmt.Printf("   ✨ %s\n", file)
	}

	fmt.Println("\n📝 Próximos pasos:")

	if projectInfo.Architecture == "modular" {
		fmt.Printf("   1. Registra el router en cmd/loom/main.go:\n")
		fmt.Printf("      %sRouter := %s.NewRouter()\n", moduleName, moduleName)
		fmt.Printf("      router.PathPrefix(\"/%s\").Handler(%sRouter)\n\n", moduleName, moduleName)
	} else {
		fmt.Printf("   1. Registra las rutas en cmd/loom/main.go:\n")
		fmt.Printf("      %sHandler := handlers.New%sHandler()\n", moduleName, toPascalCase(moduleName))
		fmt.Printf("      router.HandleFunc(\"/%s\", %sHandler.Create).Methods(\"POST\")\n\n", moduleName, moduleName)
	}

	fmt.Println("   2. Ejecuta: go mod tidy")
	fmt.Println("   3. Implementa la lógica de negocio en los archivos generados")

	return nil
}

// toPascalCase convierte una cadena a PascalCase
func toPascalCase(s string) string {
	if len(s) == 0 {
		return s
	}
	// Convertir primera letra a mayúscula
	if s[0] >= 'a' && s[0] <= 'z' {
		return string(s[0]-32) + s[1:]
	}
	return s
}
