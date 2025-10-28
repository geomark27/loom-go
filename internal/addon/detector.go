package addon

import (
	"os"
	"strings"
)

// ProjectDetector detecta qué addons están instalados en el proyecto
type ProjectDetector struct {
	projectRoot string
}

// NewProjectDetector crea un nuevo detector
func NewProjectDetector(projectRoot string) *ProjectDetector {
	return &ProjectDetector{
		projectRoot: projectRoot,
	}
}

// DetectRouter detecta qué router está usando el proyecto
func (pd *ProjectDetector) DetectRouter() string {
	goModContent, err := ReadFile("go.mod")
	if err != nil {
		return "unknown"
	}

	if strings.Contains(goModContent, "github.com/gin-gonic/gin") {
		return "gin"
	}
	if strings.Contains(goModContent, "github.com/go-chi/chi") {
		return "chi"
	}
	if strings.Contains(goModContent, "github.com/labstack/echo") {
		return "echo"
	}
	if strings.Contains(goModContent, "github.com/gorilla/mux") {
		return "gorilla-mux"
	}

	return "none"
}

// DetectORM detecta qué ORM está usando el proyecto
func (pd *ProjectDetector) DetectORM() string {
	goModContent, err := ReadFile("go.mod")
	if err != nil {
		return "unknown"
	}

	if strings.Contains(goModContent, "gorm.io/gorm") {
		return "gorm"
	}
	if strings.Contains(goModContent, "github.com/sqlc-dev/sqlc") {
		return "sqlc"
	}
	if strings.Contains(goModContent, "entgo.io/ent") {
		return "ent"
	}

	return "none"
}

// DetectDatabase detecta qué drivers de base de datos están instalados
func (pd *ProjectDetector) DetectDatabase() []string {
	databases := []string{}
	goModContent, err := ReadFile("go.mod")
	if err != nil {
		return databases
	}

	if strings.Contains(goModContent, "github.com/lib/pq") ||
		strings.Contains(goModContent, "gorm.io/driver/postgres") {
		databases = append(databases, "postgres")
	}
	if strings.Contains(goModContent, "github.com/go-sql-driver/mysql") ||
		strings.Contains(goModContent, "gorm.io/driver/mysql") {
		databases = append(databases, "mysql")
	}
	if strings.Contains(goModContent, "go.mongodb.org/mongo-driver") {
		databases = append(databases, "mongodb")
	}
	if strings.Contains(goModContent, "github.com/redis/go-redis") {
		databases = append(databases, "redis")
	}

	return databases
}

// DetectAuth detecta qué sistema de autenticación está instalado
func (pd *ProjectDetector) DetectAuth() string {
	// Verificar si existe internal/auth o pkg/auth
	if FileExists("internal/auth") || FileExists("pkg/auth") {
		// Buscar JWT
		authFiles := []string{
			"internal/auth/jwt.go",
			"pkg/auth/jwt.go",
		}

		for _, file := range authFiles {
			if FileExists(file) {
				content, _ := ReadFile(file)
				if strings.Contains(content, "github.com/golang-jwt/jwt") {
					return "jwt"
				}
			}
		}

		// Buscar OAuth2
		authFiles = []string{
			"internal/auth/oauth2.go",
			"pkg/auth/oauth2.go",
		}

		for _, file := range authFiles {
			if FileExists(file) {
				return "oauth2"
			}
		}

		return "custom"
	}

	return "none"
}

// DetectDocker detecta si el proyecto tiene Docker configurado
func (pd *ProjectDetector) DetectDocker() bool {
	return FileExists("Dockerfile") || FileExists("docker-compose.yml")
}

// GetProjectStatus retorna el estado completo del proyecto
func (pd *ProjectDetector) GetProjectStatus() map[string]interface{} {
	return map[string]interface{}{
		"router":    pd.DetectRouter(),
		"orm":       pd.DetectORM(),
		"databases": pd.DetectDatabase(),
		"auth":      pd.DetectAuth(),
		"docker":    pd.DetectDocker(),
	}
}

// HasDatabaseConfigured verifica si hay al menos una base de datos configurada
func (pd *ProjectDetector) HasDatabaseConfigured() bool {
	databases := pd.DetectDatabase()
	return len(databases) > 0
}

// GetArchitecture detecta la arquitectura del proyecto
func (pd *ProjectDetector) GetArchitecture() string {
	if _, err := os.Stat("internal/modules"); err == nil {
		return "modular"
	}
	if _, err := os.Stat("internal/app"); err == nil {
		return "layered"
	}
	return "unknown"
}
