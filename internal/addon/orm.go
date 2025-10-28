package addon

import "fmt"

// ORMAddon gestiona la instalación de ORMs
type ORMAddon struct {
	projectRoot  string
	architecture string
	ormType      string // "gorm", "sqlc"
}

// NewORMAddon crea un nuevo addon de ORM
func NewORMAddon(projectRoot, architecture, ormType string) *ORMAddon {
	return &ORMAddon{
		projectRoot:  projectRoot,
		architecture: architecture,
		ormType:      ormType,
	}
}

func (o *ORMAddon) Name() string {
	return fmt.Sprintf("ORM %s", o.ormType)
}

func (o *ORMAddon) Description() string {
	descriptions := map[string]string{
		"gorm": "ORM completo con features avanzadas",
		"sqlc": "Generador de código type-safe desde SQL",
	}
	return descriptions[o.ormType]
}

func (o *ORMAddon) IsInstalled() (bool, error) {
	detector := NewProjectDetector(o.projectRoot)
	currentORM := detector.DetectORM()
	return currentORM == o.ormType, nil
}

func (o *ORMAddon) CanInstall() (bool, string, error) {
	// Verificar que no haya otro ORM instalado
	detector := NewProjectDetector(o.projectRoot)
	currentORM := detector.DetectORM()

	if currentORM != "none" && currentORM != o.ormType {
		return false, fmt.Sprintf("Ya tienes %s instalado. Usa --force para reemplazar", currentORM), nil
	}

	return true, "", nil
}

func (o *ORMAddon) GetConflicts() []string {
	conflicts := []string{"gorm", "sqlc", "ent"}
	filtered := []string{}
	for _, c := range conflicts {
		if c != o.ormType {
			filtered = append(filtered, c)
		}
	}
	return filtered
}

func (o *ORMAddon) Install(force bool) error {
	switch o.ormType {
	case "gorm":
		return o.installGORM()
	case "sqlc":
		return o.installSQLC()
	default:
		return fmt.Errorf("ORM no soportado: %s", o.ormType)
	}
}

func (o *ORMAddon) installGORM() error {
	fmt.Println("   📦 Instalando GORM...")

	// Añadir dependencia
	if err := UpdateGoMod("gorm.io/gorm", "v1.25.5"); err != nil {
		return err
	}

	// Crear estructura de database
	fmt.Println("   📁 Creando internal/database/...")

	// TODO: Implementar creación completa de archivos GORM
	fmt.Println("   ⚠️  Implementación completa próximamente")
	fmt.Println("   💡 Ejecuta 'go get gorm.io/gorm' manualmente por ahora")

	return nil
}

func (o *ORMAddon) installSQLC() error {
	fmt.Println("   📦 Instalando sqlc...")

	// TODO: Implementar instalación de sqlc
	fmt.Println("   ⚠️  Implementación completa próximamente")
	fmt.Println("   💡 Visita https://docs.sqlc.dev/en/latest/overview/install.html")

	return nil
}
