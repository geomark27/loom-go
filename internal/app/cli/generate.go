package cli

import (
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Genera componentes en un proyecto existente",
	Long: `Genera módulos, handlers, services, models y otros componentes
en un proyecto Loom existente.

El comando detectará automáticamente si tu proyecto usa arquitectura
Layered o Modular y generará el código apropiado.

Ejemplos:
  loom generate module products
  loom generate handler orders
  loom generate service email
  loom generate model Category
  loom generate middleware auth`,
	Aliases: []string{"gen", "g"},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Flags globales para todos los subcomandos de generate
	generateCmd.PersistentFlags().Bool("force", false, "Sobrescribir archivos existentes")
	generateCmd.PersistentFlags().Bool("dry-run", false, "Mostrar qué se generaría sin crear archivos")
}
