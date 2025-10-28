package addon

import (
	"fmt"
	"path/filepath"
)

// DockerAddon gestiona la configuración de Docker
type DockerAddon struct {
	projectRoot  string
	architecture string
}

// NewDockerAddon crea un nuevo addon de Docker
func NewDockerAddon(projectRoot, architecture string) *DockerAddon {
	return &DockerAddon{
		projectRoot:  projectRoot,
		architecture: architecture,
	}
}

func (d *DockerAddon) Name() string {
	return "Docker"
}

func (d *DockerAddon) Description() string {
	return "Containerización con Docker y Docker Compose"
}

func (d *DockerAddon) IsInstalled() (bool, error) {
	detector := NewProjectDetector(d.projectRoot)
	return detector.DetectDocker(), nil
}

func (d *DockerAddon) CanInstall() (bool, string, error) {
	// Docker siempre se puede instalar
	return true, "", nil
}

func (d *DockerAddon) GetConflicts() []string {
	return []string{} // Docker no tiene conflictos
}

func (d *DockerAddon) Install(force bool) error {
	// 1. Crear Dockerfile
	if err := d.createDockerfile(); err != nil {
		return fmt.Errorf("error al crear Dockerfile: %w", err)
	}

	// 2. Crear .dockerignore
	if err := d.createDockerignore(); err != nil {
		return fmt.Errorf("error al crear .dockerignore: %w", err)
	}

	// 3. Crear docker-compose.yml
	if err := d.createDockerCompose(); err != nil {
		return fmt.Errorf("error al crear docker-compose.yml: %w", err)
	}

	// 4. Actualizar Makefile si existe
	if FileExists("Makefile") {
		if err := d.updateMakefile(); err != nil {
			return fmt.Errorf("error al actualizar Makefile: %w", err)
		}
	}

	fmt.Println("\n📝 Archivos Docker creados:")
	fmt.Println("   ✨ Dockerfile")
	fmt.Println("   ✨ .dockerignore")
	fmt.Println("   ✨ docker-compose.yml")

	return nil
}

func (d *DockerAddon) createDockerfile() error {
	fmt.Println("   📝 Creando Dockerfile...")

	content := `# Build stage
FROM golang:1.23-alpine AS builder

# Instalar dependencias de build
RUN apk add --no-cache git

WORKDIR /app

# Copiar go.mod y go.sum
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar código fuente
COPY . .

# Build de la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/app

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar el binario desde el builder
COPY --from=builder /app/main .
COPY --from=builder /app/.env.example .env

# Exponer el puerto
EXPOSE 8080

# Comando para ejecutar
CMD ["./main"]
`

	return WriteFile(filepath.Join(d.projectRoot, "Dockerfile"), content)
}

func (d *DockerAddon) createDockerignore() error {
	fmt.Println("   📝 Creando .dockerignore...")

	content := `# Git
.git
.gitignore

# Archivos de desarrollo
.env
*.log
*.exe
*.test
*.out

# Directorios
tmp/
vendor/
.vscode/
.idea/

# Documentación
*.md
docs/

# CI/CD
.github/
.gitlab-ci.yml

# Docker
Dockerfile
.dockerignore
docker-compose*.yml

# Backups
.loom-backups/
`

	return WriteFile(filepath.Join(d.projectRoot, ".dockerignore"), content)
}

func (d *DockerAddon) createDockerCompose() error {
	fmt.Println("   📝 Creando docker-compose.yml...")

	// Detectar si tiene base de datos configurada
	detector := NewProjectDetector(d.projectRoot)
	databases := detector.DetectDatabase()

	content := `version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - PORT=8080
`

	// Si tiene PostgreSQL, añadirlo
	hasPostgres := false
	for _, db := range databases {
		if db == "postgres" {
			hasPostgres = true
			break
		}
	}

	if hasPostgres {
		content += `      - DB_HOST=postgres
      - DB_PORT=5432
      - DB_USER=postgres
      - DB_PASSWORD=postgres
      - DB_NAME=app_db
      - DB_SSLMODE=disable
    depends_on:
      - postgres
    volumes:
      - .:/app
    networks:
      - app-network

  postgres:
    image: postgres:16-alpine
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: app_db
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    networks:
      - app-network

volumes:
  postgres_data:

networks:
  app-network:
    driver: bridge
`
	} else {
		content += `    volumes:
      - .:/app
    networks:
      - app-network

networks:
  app-network:
    driver: bridge
`
	}

	return WriteFile(filepath.Join(d.projectRoot, "docker-compose.yml"), content)
}

func (d *DockerAddon) updateMakefile() error {
	fmt.Println("   📝 Actualizando Makefile...")

	content, err := ReadFile("Makefile")
	if err != nil {
		return err
	}

	// Verificar si ya tiene comandos Docker
	if content != "" && content[len(content)-1:] != "\n" {
		content += "\n"
	}

	dockerCommands := `
# Docker commands
.PHONY: docker-build docker-up docker-down docker-logs

docker-build:
	@echo "Building Docker image..."
	docker-compose build

docker-up:
	@echo "Starting containers..."
	docker-compose up -d

docker-down:
	@echo "Stopping containers..."
	docker-compose down

docker-logs:
	@echo "Showing logs..."
	docker-compose logs -f app

docker-clean:
	@echo "Cleaning Docker resources..."
	docker-compose down -v
	docker system prune -f
`

	content += dockerCommands

	return WriteFile("Makefile", content)
}
