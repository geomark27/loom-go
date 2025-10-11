package cli

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/geomark27/loom-go/internal/generator"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [nombre-del-proyecto]",
	Short: "Crea un nuevo proyecto Go con Loom",
	Long: `Crea un nuevo proyecto Go siguiendo las mejores prácticas y la estructura 
estándar de golang-standard/project-layout.

El proyecto generado incluirá:
- Estructura de directorios idiomática
- go.mod configurado
- Servidor web básico con net/http
- README.md con instrucciones`,
	Args: cobra.ExactArgs(1),
	RunE: runNewCommand,
}

func runNewCommand(cmd *cobra.Command, args []string) error {
	projectName := args[0]

	// Validar nombre del proyecto
	if err := validateProjectName(projectName); err != nil {
		return fmt.Errorf("nombre de proyecto inválido: %w", err)
	}

	// Obtener el directorio actual como directorio base
	baseDir := "."
	projectPath := filepath.Join(baseDir, projectName)

	// Crear la configuración del proyecto
	config := &generator.ProjectConfig{
		Name:        projectName,
		Path:        projectPath,
		ModuleName:  fmt.Sprintf("github.com/tu-usuario/%s", projectName), // TODO: hacer configurable
		Description: fmt.Sprintf("Proyecto %s generado con Loom", projectName),
	}

	// Generar el proyecto
	gen := generator.New()
	if err := gen.GenerateProject(config); err != nil {
		return fmt.Errorf("error generando proyecto: %w", err)
	}

	// Mensaje de éxito
	fmt.Printf("✅ Proyecto '%s' creado exitosamente en %s\n", projectName, projectPath)
	fmt.Printf("\nPróximos pasos:\n")
	fmt.Printf("  cd %s\n", projectName)
	fmt.Printf("  go mod tidy\n")
	fmt.Printf("  go run cmd/%s/main.go\n", projectName)

	return nil
}

func validateProjectName(name string) error {
	if name == "" {
		return fmt.Errorf("el nombre no puede estar vacío")
	}

	if strings.Contains(name, " ") {
		return fmt.Errorf("el nombre no puede contener espacios")
	}

	// Verificar caracteres válidos para nombres de directorios
	if strings.ContainsAny(name, `<>:"/\|?*`) {
		return fmt.Errorf("el nombre contiene caracteres no válidos")
	}

	return nil
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Flags específicos del comando new
	newCmd.Flags().StringP("module", "m", "", "Nombre del módulo Go (por defecto: github.com/tu-usuario/nombre-proyecto)")
}
