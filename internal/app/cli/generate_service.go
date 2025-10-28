package cli

import (
	"fmt"

	"github.com/geomark27/loom-go/internal/generator"
	"github.com/spf13/cobra"
)

var generateServiceCmd = &cobra.Command{
	Use:   "service [nombre]",
	Short: "Genera un service con lógica de negocio",
	Long: `Genera un archivo service para contener la lógica de negocio.

El archivo se generará en la ubicación apropiada según la arquitectura:
  - Layered: internal/app/services/{nombre}_service.go
  - Modular: internal/modules/{nombre}/service.go

Ejemplos:
  loom generate service products
  loom generate service email --force`,
	Aliases: []string{"svc", "s"},
	Args:    cobra.ExactArgs(1),
	RunE:    runGenerateService,
}

func init() {
	generateCmd.AddCommand(generateServiceCmd)
}

func runGenerateService(cmd *cobra.Command, args []string) error {
	name := args[0]
	force, _ := cmd.Flags().GetBool("force")
	dryRun, _ := cmd.Flags().GetBool("dry-run")

	projectInfo, err := generator.DetectProject()
	if err != nil {
		return fmt.Errorf("error: no se detectó un proyecto Loom válido. %w", err)
	}

	if err := generator.ValidateComponentName(name); err != nil {
		return fmt.Errorf("nombre de service inválido: %w", err)
	}

	fmt.Printf("🔍 Proyecto: %s (%s)\n", projectInfo.Name, projectInfo.Architecture)
	fmt.Printf("📦 Generando service: %s\n\n", name)

	gen := generator.NewModuleGenerator(projectInfo)
	files, err := gen.GenerateService(name, force, dryRun)
	if err != nil {
		return fmt.Errorf("error al generar service: %w", err)
	}

	if dryRun {
		fmt.Println("📋 Archivo que se generaría:")
		for _, file := range files {
			fmt.Printf("   ✨ %s\n", file)
		}
		fmt.Println("\n💡 Ejecuta sin --dry-run para crear el archivo")
		return nil
	}

	fmt.Println("✅ Service generado exitosamente!")
	fmt.Println("\n📝 Archivo creado:")
	for _, file := range files {
		fmt.Printf("   ✨ %s\n", file)
	}

	return nil
}
