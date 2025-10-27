package cli

import (
	"fmt"
	"os/exec"
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
	modular    bool
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
		// Intentar detectar el usuario de GitHub desde git config
		githubUser := detectGitHubUser()
		if githubUser != "" {
			module = fmt.Sprintf("github.com/%s/%s", githubUser, projectName)
		} else {
			// Fallback: usar el nombre del proyecto directamente
			module = projectName
		}
	}

	// Determinar arquitectura
	architecture := "layered"
	if modular {
		architecture = "modular"
	}

	// Crear la configuración del proyecto
	config := &generator.ProjectConfig{
		Name:         projectName,
		Path:         projectPath,
		ModuleName:   module,
		Description:  fmt.Sprintf("Proyecto %s generado con Loom", projectName),
		UseHelpers:   !standalone, // UseHelpers es true por defecto, false si --standalone está activo
		IsModular:    modular,
		Architecture: architecture,
	}

	// Generar el proyecto
	gen := generator.New()
	if err := gen.GenerateProject(config); err != nil {
		return fmt.Errorf("error generando proyecto: %w", err)
	}

	// Mensaje de éxito con información de arquitectura
	fmt.Printf("✅ Proyecto '%s' creado exitosamente en %s\n", projectName, projectPath)

	// Información de arquitectura
	if config.IsModular {
		fmt.Printf("\n🏗️  Arquitectura: Modular (por dominios)\n")
		fmt.Printf("   → Ideal para: Proyectos grandes (20+ endpoints), equipos, microservicios\n")
		fmt.Printf("   → Módulos: users (ejemplo generado)\n")
		fmt.Printf("\n💡 Tips:\n")
		fmt.Printf("   • Usa 'loom generate module <name>' para agregar módulos\n")
		fmt.Printf("   • Mantén módulos independientes (usa Event Bus para comunicación)\n")
		fmt.Printf("   • Cada módulo tiene su propio ports.go con interfaces\n")
	} else {
		fmt.Printf("\n🏗️  Arquitectura: Layered (por capas)\n")
		fmt.Printf("   → Ideal para: APIs pequeñas (< 20 endpoints), MVPs, prototipos\n")
		fmt.Printf("   → Estructura: handlers → services → repositories\n")
		fmt.Printf("\n💡 Tips:\n")
		fmt.Printf("   • Empieza simple, escala cuando lo necesites\n")
		fmt.Printf("   • Usa 'loom generate module <name>' para agregar recursos\n")
		fmt.Printf("   • Considera --modular si tienes 20+ endpoints\n")
	}

	// Información de helpers
	if config.UseHelpers {
		fmt.Printf("\n📦 Helpers: Incluidos (validación, respuestas, logging)\n")
	} else {
		fmt.Printf("\n🔧 Modo: Standalone (sin dependencias externas)\n")
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

// detectGitHubUser intenta detectar el usuario de GitHub desde la configuración de git
func detectGitHubUser() string {
	// Intentar obtener github.user
	cmd := exec.Command("git", "config", "github.user")
	if output, err := cmd.Output(); err == nil {
		user := strings.TrimSpace(string(output))
		if user != "" {
			return user
		}
	}

	// Fallback: intentar extraer de la URL del remote origin
	cmd = exec.Command("git", "config", "remote.origin.url")
	if output, err := cmd.Output(); err == nil {
		url := strings.TrimSpace(string(output))
		// Parsear URLs como: git@github.com:username/repo.git o https://github.com/username/repo.git
		if strings.Contains(url, "github.com") {
			// Para SSH: git@github.com:username/repo.git
			if strings.HasPrefix(url, "git@github.com:") {
				parts := strings.Split(strings.TrimPrefix(url, "git@github.com:"), "/")
				if len(parts) > 0 {
					return parts[0]
				}
			}
			// Para HTTPS: https://github.com/username/repo.git
			if strings.Contains(url, "github.com/") {
				parts := strings.Split(url, "github.com/")
				if len(parts) > 1 {
					userRepo := strings.Split(parts[1], "/")
					if len(userRepo) > 0 {
						return userRepo[0]
					}
				}
			}
		}
	}

	return ""
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Flags específicos del comando new
	newCmd.Flags().StringVarP(&moduleName, "module", "m", "", "Nombre del módulo Go (detecta automáticamente desde git config o usa el nombre del proyecto)")
	newCmd.Flags().BoolVar(&standalone, "standalone", false, "Generar proyecto sin helpers de Loom (código 100% independiente)")
	newCmd.Flags().BoolVar(&modular, "modular", false, "Generar arquitectura modular por dominio (recomendado para proyectos grandes con 20+ endpoints)")
}
