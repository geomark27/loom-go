package cli

import (
	"fmt"

	"github.com/geomark27/loom-go/internal/generator"
	"github.com/spf13/cobra"
)

var generateMiddlewareCmd = &cobra.Command{
	Use:   "middleware [nombre]",
	Short: "Genera un middleware HTTP",
	Long: `Genera un middleware para interceptar peticiones HTTP.

El archivo se generará en la ubicación apropiada según la arquitectura:
  - Layered: internal/app/middleware/{nombre}.go
  - Modular: internal/middleware/{nombre}.go

Ejemplos:
  loom generate middleware auth
  loom generate middleware logger --force`,
	Aliases: []string{"mw"},
	Args:    cobra.ExactArgs(1),
	RunE:    runGenerateMiddleware,
}

func init() {
	generateCmd.AddCommand(generateMiddlewareCmd)
}

func runGenerateMiddleware(cmd *cobra.Command, args []string) error {
	name := args[0]
	force, _ := cmd.Flags().GetBool("force")
	dryRun, _ := cmd.Flags().GetBool("dry-run")

	projectInfo, err := generator.DetectProject()
	if err != nil {
		return fmt.Errorf("error: no se detectó un proyecto Loom válido. %w", err)
	}

	if err := generator.ValidateComponentName(name); err != nil {
		return fmt.Errorf("nombre de middleware inválido: %w", err)
	}

	fmt.Printf("🔍 Proyecto: %s (%s)\n", projectInfo.Name, projectInfo.Architecture)
	fmt.Printf("📦 Generando middleware: %s\n\n", name)

	gen := generator.NewModuleGenerator(projectInfo)
	files, err := gen.GenerateMiddleware(name, force, dryRun)
	if err != nil {
		return fmt.Errorf("error al generar middleware: %w", err)
	}

	if dryRun {
		fmt.Println("📋 Archivo que se generaría:")
		for _, file := range files {
			fmt.Printf("   ✨ %s\n", file)
		}
		fmt.Println("\n💡 Ejecuta sin --dry-run para crear el archivo")
		return nil
	}

	fmt.Println("✅ Middleware generado exitosamente!")
	fmt.Println("\n📝 Archivo creado:")
	for _, file := range files {
		fmt.Printf("   ✨ %s\n", file)
	}

	fmt.Println("\n📝 Próximo paso:")
	fmt.Println("   Registra el middleware en tu router o en rutas específicas")

	return nil
}
