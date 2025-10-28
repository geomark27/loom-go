package addon

import "fmt"

// AuthAddon gestiona sistemas de autenticación
type AuthAddon struct {
	projectRoot  string
	architecture string
	authType     string // "jwt", "oauth2"
}

// NewAuthAddon crea un nuevo addon de autenticación
func NewAuthAddon(projectRoot, architecture, authType string) *AuthAddon {
	return &AuthAddon{
		projectRoot:  projectRoot,
		architecture: architecture,
		authType:     authType,
	}
}

func (a *AuthAddon) Name() string {
	return fmt.Sprintf("Auth %s", a.authType)
}

func (a *AuthAddon) Description() string {
	descriptions := map[string]string{
		"jwt":    "JSON Web Tokens para autenticación stateless",
		"oauth2": "OAuth 2.0 para autenticación con terceros",
	}
	return descriptions[a.authType]
}

func (a *AuthAddon) IsInstalled() (bool, error) {
	detector := NewProjectDetector(a.projectRoot)
	currentAuth := detector.DetectAuth()
	return currentAuth == a.authType, nil
}

func (a *AuthAddon) CanInstall() (bool, string, error) {
	// Verificar que no haya otro sistema de auth
	detector := NewProjectDetector(a.projectRoot)
	currentAuth := detector.DetectAuth()

	if currentAuth != "none" && currentAuth != a.authType {
		return false, fmt.Sprintf("Ya tienes %s instalado. Usa --force para reemplazar", currentAuth), nil
	}

	return true, "", nil
}

func (a *AuthAddon) GetConflicts() []string {
	conflicts := []string{"jwt", "oauth2", "session"}
	filtered := []string{}
	for _, c := range conflicts {
		if c != a.authType {
			filtered = append(filtered, c)
		}
	}
	return filtered
}

func (a *AuthAddon) Install(force bool) error {
	switch a.authType {
	case "jwt":
		return a.installJWT()
	case "oauth2":
		return a.installOAuth2()
	default:
		return fmt.Errorf("sistema de auth no soportado: %s", a.authType)
	}
}

func (a *AuthAddon) installJWT() error {
	fmt.Println("   📦 Instalando JWT Auth...")

	// Añadir dependencia
	if err := UpdateGoMod("github.com/golang-jwt/jwt/v5", "v5.2.0"); err != nil {
		return err
	}

	// Actualizar .env.example
	envVars := map[string]string{
		"JWT_SECRET":     "your-secret-key-change-this-in-production",
		"JWT_EXPIRATION": "24h",
	}

	if err := UpdateEnvExample(envVars, "JWT Authentication"); err != nil {
		return err
	}

	fmt.Println("   ✅ JWT configurado")
	fmt.Println("   💡 Próximamente: Generación automática de estructura de auth")

	// TODO: Crear estructura completa de auth
	// - internal/auth/jwt.go
	// - internal/auth/middleware.go
	// - internal/handlers/auth_handler.go

	return nil
}

func (a *AuthAddon) installOAuth2() error {
	fmt.Println("   📦 Instalando OAuth2...")

	// TODO: Implementar OAuth2
	fmt.Println("   ⚠️  Implementación completa próximamente")

	return nil
}
