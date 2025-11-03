package addon

import "fmt"

// AuthAddon manages authentication systems
type AuthAddon struct {
	projectRoot  string
	architecture string
	authType     string // "jwt", "oauth2"
}

// NewAuthAddon creates a new authentication addon
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
		"jwt":    "JSON Web Tokens for stateless authentication",
		"oauth2": "OAuth 2.0 for third-party authentication",
	}
	return descriptions[a.authType]
}

func (a *AuthAddon) IsInstalled() (bool, error) {
	detector := NewProjectDetector(a.projectRoot)
	currentAuth := detector.DetectAuth()
	return currentAuth == a.authType, nil
}

func (a *AuthAddon) CanInstall() (bool, string, error) {
	// Check that there's no other auth system
	detector := NewProjectDetector(a.projectRoot)
	currentAuth := detector.DetectAuth()

	if currentAuth != "none" && currentAuth != a.authType {
		return false, fmt.Sprintf("You already have %s installed. Use --force to replace", currentAuth), nil
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
		return fmt.Errorf("unsupported auth system: %s", a.authType)
	}
}

func (a *AuthAddon) installJWT() error {
	fmt.Println("   📦 Installing JWT Auth...")

	// Add dependency
	if err := UpdateGoMod("github.com/golang-jwt/jwt/v5", "v5.2.0"); err != nil {
		return err
	}

	// Update .env.example
	envVars := map[string]string{
		"JWT_SECRET":     "your-secret-key-change-this-in-production",
		"JWT_EXPIRATION": "24h",
	}

	if err := UpdateEnvExample(envVars, "JWT Authentication"); err != nil {
		return err
	}

	fmt.Println("   ✅ JWT configured")
	fmt.Println("   💡 Coming soon: Automatic generation of auth structure")

	// TODO: Create complete auth structure
	// - internal/auth/jwt.go
	// - internal/auth/middleware.go
	// - internal/handlers/auth_handler.go

	return nil
}

func (a *AuthAddon) installOAuth2() error {
	fmt.Println("   📦 Installing OAuth2...")

	// TODO: Implement OAuth2
	fmt.Println("   ⚠️  Full implementation coming soon")

	return nil
}
