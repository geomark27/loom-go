package addon

import "fmt"

// ORMAddon manages ORM installation
type ORMAddon struct {
	projectRoot  string
	architecture string
	ormType      string // "gorm", "sqlc"
}

// NewORMAddon creates a new ORM addon
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
		"gorm": "Complete ORM with advanced features",
		"sqlc": "Type-safe code generator from SQL",
	}
	return descriptions[o.ormType]
}

func (o *ORMAddon) IsInstalled() (bool, error) {
	detector := NewProjectDetector(o.projectRoot)
	currentORM := detector.DetectORM()
	return currentORM == o.ormType, nil
}

func (o *ORMAddon) CanInstall() (bool, string, error) {
	// Check that there's no other ORM installed
	detector := NewProjectDetector(o.projectRoot)
	currentORM := detector.DetectORM()

	if currentORM != "none" && currentORM != o.ormType {
		return false, fmt.Sprintf("You already have %s installed. Use --force to replace", currentORM), nil
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
		return fmt.Errorf("unsupported ORM: %s", o.ormType)
	}
}

func (o *ORMAddon) installGORM() error {
	fmt.Println("   📦 Installing GORM...")

	// Add dependency
	if err := UpdateGoMod("gorm.io/gorm", "v1.25.5"); err != nil {
		return err
	}

	// Create database structure
	fmt.Println("   📁 Creating internal/database/...")

	// TODO: Implement complete GORM file creation
	fmt.Println("   ⚠️  Full implementation coming soon")
	fmt.Println("   💡 Run 'go get gorm.io/gorm' manually for now")

	return nil
}

func (o *ORMAddon) installSQLC() error {
	fmt.Println("   📦 Installing sqlc...")

	// TODO: Implement sqlc installation
	fmt.Println("   ⚠️  Full implementation coming soon")
	fmt.Println("   💡 Visit https://docs.sqlc.dev/en/latest/overview/install.html")

	return nil
}
