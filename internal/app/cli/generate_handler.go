package cli

import (
	"fmt"

	"github.com/geomark27/loom-go/internal/generator"
	"github.com/spf13/cobra"
)

var generateHandlerCmd = &cobra.Command{
	Use:   "handler [nombre]",
	Short: "Genera un handler HTTP",
	Long: `Genera un archivo handler para manejar peticiones HTTP.

El archivo se generará en la ubicación apropiada según la arquitectura:
  - Layered: internal/app/handlers/{nombre}_handler.go
  - Modular: internal/modules/{nombre}/handler.go

Ejemplos:
  loom generate handler products
  loom generate handler users --force`,
	Aliases: []string{"h"},
	Args:    cobra.ExactArgs(1),
	RunE:    runGenerateHandler,
}

func init() {
	generateCmd.AddCommand(generateHandlerCmd)
}

func runGenerateHandler(cmd *cobra.Command, args []string) error {
	name := args[0]
	force, _ := cmd.Flags().GetBool("force")
	dryRun, _ := cmd.Flags().GetBool("dry-run")

	projectInfo, err := generator.DetectProject()
	if err != nil {
		return fmt.Errorf("error: no se detectó un proyecto Loom válido. %w", err)
	}

	if err := generator.ValidateComponentName(name); err != nil {
		return fmt.Errorf("nombre de handler inválido: %w", err)
	}

	fmt.Printf("🔍 Proyecto: %s (%s)\n", projectInfo.Name, projectInfo.Architecture)
	fmt.Printf("📦 Generando handler: %s\n\n", name)

	gen := generator.NewModuleGenerator(projectInfo)
	files, err := gen.GenerateHandler(name, force, dryRun)
	if err != nil {
		return fmt.Errorf("error al generar handler: %w", err)
	}

	if dryRun {
		fmt.Println("📋 Archivo que se generaría:")
		for _, file := range files {
			fmt.Printf("   ✨ %s\n", file)
		}
		fmt.Println("\n💡 Ejecuta sin --dry-run para crear el archivo")
		return nil
	}

	fmt.Println("✅ Handler generado exitosamente!")
	fmt.Println("\n📝 Archivo creado:")
	for _, file := range files {
		fmt.Printf("   ✨ %s\n", file)
	}

	return nil
}
