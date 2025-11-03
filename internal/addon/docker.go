package addon

import (
	"fmt"
	"path/filepath"
)

// DockerAddon manages Docker configuration
type DockerAddon struct {
	projectRoot  string
	architecture string
}

// NewDockerAddon creates a new Docker addon
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
	return "Containerization with Docker and Docker Compose"
}

func (d *DockerAddon) IsInstalled() (bool, error) {
	detector := NewProjectDetector(d.projectRoot)
	return detector.DetectDocker(), nil
}

func (d *DockerAddon) CanInstall() (bool, string, error) {
	// Docker can always be installed
	return true, "", nil
}

func (d *DockerAddon) GetConflicts() []string {
	return []string{} // Docker has no conflicts
}

func (d *DockerAddon) Install(force bool) error {
	// 1. Create Dockerfile
	if err := d.createDockerfile(); err != nil {
		return fmt.Errorf("error creating Dockerfile: %w", err)
	}

	// 2. Create .dockerignore
	if err := d.createDockerignore(); err != nil {
		return fmt.Errorf("error creating .dockerignore: %w", err)
	}

	// 3. Create docker-compose.yml
	if err := d.createDockerCompose(); err != nil {
		return fmt.Errorf("error creating docker-compose.yml: %w", err)
	}

	// 4. Update Makefile if it exists
	if FileExists("Makefile") {
		if err := d.updateMakefile(); err != nil {
			return fmt.Errorf("error updating Makefile: %w", err)
		}
	}

	fmt.Println("\n📝 Docker files created:")
	fmt.Println("   ✨ Dockerfile")
	fmt.Println("   ✨ .dockerignore")
	fmt.Println("   ✨ docker-compose.yml")

	return nil
}

func (d *DockerAddon) createDockerfile() error {
	fmt.Println("   📝 Creating Dockerfile...")

	content := `# Build stage
FROM golang:1.23-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/app

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /app/main .
COPY --from=builder /app/.env.example .env

# Expose the port
EXPOSE 8080

# Command to execute
CMD ["./main"]
`

	return WriteFile(filepath.Join(d.projectRoot, "Dockerfile"), content)
}

func (d *DockerAddon) createDockerignore() error {
	fmt.Println("   📝 Creating .dockerignore...")

	content := `# Git
.git
.gitignore

# Development files
.env
*.log
*.exe
*.test
*.out

# Directories
tmp/
vendor/
.vscode/
.idea/

# Documentation
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
	fmt.Println("   📝 Creating docker-compose.yml...")

	// Detect if a database is configured
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

	// If it has PostgreSQL, add it
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
	fmt.Println("   📝 Updating Makefile...")

	content, err := ReadFile("Makefile")
	if err != nil {
		return err
	}

	// Check if it already has Docker commands
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
