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

var (
	standalone bool
	moduleName string
)

func runNewCommand(cmd *cobra.Command, args []string) error {
	projectName := args[0]

	// Validar nombre del proyecto
	if err := validateProjectName(projectName); err != nil {
		return fmt.Errorf("nombre de proyecto inválido: %w", err)
	}

	// Obtener el directorio actual como directorio base
	baseDir := "."
	projectPath := filepath.Join(baseDir, projectName)

	// Determinar el nombre del módulo
	module := moduleName
	if module == "" {
		module = fmt.Sprintf("github.com/tu-usuario/%s", projectName)
	}

	// Crear la configuración del proyecto
	config := &generator.ProjectConfig{
		Name:        projectName,
		Path:        projectPath,
		ModuleName:  module,
		Description: fmt.Sprintf("Proyecto %s generado con Loom", projectName),
		UseHelpers:  !standalone, // UseHelpers es true por defecto, false si --standalone está activo
	}

	// Generar el proyecto
	gen := generator.New()
	if err := gen.GenerateProject(config); err != nil {
		return fmt.Errorf("error generando proyecto: %w", err)
	}

	// Mensaje de éxito
	fmt.Printf("✅ Proyecto '%s' creado exitosamente en %s\n", projectName, projectPath)

	if config.UseHelpers {
		fmt.Printf("📦 Incluye helpers de Loom para desarrollo rápido\n")
	} else {
		fmt.Printf("🔧 Proyecto standalone (sin dependencias de Loom)\n")
	}

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
	newCmd.Flags().StringVarP(&moduleName, "module", "m", "", "Nombre del módulo Go (por defecto: github.com/tu-usuario/nombre-proyecto)")
	newCmd.Flags().BoolVar(&standalone, "standalone", false, "Generar proyecto sin helpers de Loom (código 100% independiente)")
}
