package cli

import (
	"fmt"

	"github.com/geomark27/loom-go/internal/generator"
	"github.com/spf13/cobra"
)

var generateModelCmd = &cobra.Command{
	Use:   "model [nombre]",
	Short: "Genera un modelo de datos",
	Long: `Genera un archivo model con la estructura de datos.

El archivo se generará en la ubicación apropiada según la arquitectura:
  - Layered: internal/app/models/{nombre}.go
  - Modular: internal/modules/{nombre}/model.go

Ejemplos:
  loom generate model Product
  loom generate model User --force`,
	Aliases: []string{"mod"},
	Args:    cobra.ExactArgs(1),
	RunE:    runGenerateModel,
}

func init() {
	generateCmd.AddCommand(generateModelCmd)
}

func runGenerateModel(cmd *cobra.Command, args []string) error {
	name := args[0]
	force, _ := cmd.Flags().GetBool("force")
	dryRun, _ := cmd.Flags().GetBool("dry-run")

	projectInfo, err := generator.DetectProject()
	if err != nil {
		return fmt.Errorf("error: no se detectó un proyecto Loom válido. %w", err)
	}

	if err := generator.ValidateComponentName(name); err != nil {
		return fmt.Errorf("nombre de model inválido: %w", err)
	}

	fmt.Printf("🔍 Proyecto: %s (%s)\n", projectInfo.Name, projectInfo.Architecture)
	fmt.Printf("📦 Generando model: %s\n\n", name)

	gen := generator.NewModuleGenerator(projectInfo)
	files, err := gen.GenerateModel(name, force, dryRun)
	if err != nil {
		return fmt.Errorf("error al generar model: %w", err)
	}

	if dryRun {
		fmt.Println("📋 Archivo que se generaría:")
		for _, file := range files {
			fmt.Printf("   ✨ %s\n", file)
		}
		fmt.Println("\n💡 Ejecuta sin --dry-run para crear el archivo")
		return nil
	}

	fmt.Println("✅ Model generado exitosamente!")
	fmt.Println("\n📝 Archivo creado:")
	for _, file := range files {
		fmt.Printf("   ✨ %s\n", file)
	}

	return nil
}
