package addon

import (
	"fmt"
	"path/filepath"
	"strings"
)

// RouterAddon gestiona la instalación de routers HTTP
type RouterAddon struct {
	projectRoot  string
	architecture string
	routerType   string // "gin", "chi", "echo"
}

// NewRouterAddon crea un nuevo addon de router
func NewRouterAddon(projectRoot, architecture, routerType string) *RouterAddon {
	return &RouterAddon{
		projectRoot:  projectRoot,
		architecture: architecture,
		routerType:   routerType,
	}
}

func (r *RouterAddon) Name() string {
	return fmt.Sprintf("Router %s", strings.ToUpper(r.routerType))
}

func (r *RouterAddon) Description() string {
	descriptions := map[string]string{
		"gin":  "Framework web rápido con excelente performance",
		"chi":  "Router ligero y composable compatible con net/http",
		"echo": "Framework minimalista de alto rendimiento",
	}
	return descriptions[r.routerType]
}

func (r *RouterAddon) IsInstalled() (bool, error) {
	detector := NewProjectDetector(r.projectRoot)
	currentRouter := detector.DetectRouter()
	return currentRouter == r.routerType, nil
}

func (r *RouterAddon) CanInstall() (bool, string, error) {
	// Los routers siempre se pueden instalar (reemplazan el existente)
	return true, "", nil
}

func (r *RouterAddon) GetConflicts() []string {
	// Todos los routers son conflictivos entre sí
	conflicts := []string{"gin", "chi", "echo", "gorilla-mux"}

	// Remover el router actual de los conflictos
	filtered := []string{}
	for _, c := range conflicts {
		if c != r.routerType {
			filtered = append(filtered, c)
		}
	}

	return filtered
}

func (r *RouterAddon) Install(force bool) error {
	// 1. Actualizar go.mod con la dependencia
	if err := r.addDependency(); err != nil {
		return fmt.Errorf("error al añadir dependencia: %w", err)
	}

	// 2. Actualizar archivos de servidor
	if err := r.updateServerFiles(); err != nil {
		return fmt.Errorf("error al actualizar servidor: %w", err)
	}

	// 3. Actualizar handlers (si es necesario)
	fmt.Println("⚠️  Nota: Los handlers existentes necesitarán ser actualizados manualmente")
	fmt.Println("   para usar la nueva API del router")

	return nil
}

func (r *RouterAddon) addDependency() error {
	modules := map[string]string{
		"gin":  "github.com/gin-gonic/gin v1.9.1",
		"chi":  "github.com/go-chi/chi/v5 v5.0.10",
		"echo": "github.com/labstack/echo/v4 v4.11.3",
	}

	parts := strings.Split(modules[r.routerType], " ")
	if len(parts) != 2 {
		return fmt.Errorf("formato de módulo inválido")
	}

	fmt.Printf("   📦 Añadiendo dependencia: %s\n", modules[r.routerType])
	return UpdateGoMod(parts[0], parts[1])
}

func (r *RouterAddon) updateServerFiles() error {
	serverPath := r.getServerPath()

	fmt.Printf("   📝 Actualizando %s\n", serverPath)

	// Generar nuevo contenido según el router
	newContent := r.generateServerContent()

	return WriteFile(serverPath, newContent)
}

func (r *RouterAddon) getServerPath() string {
	if r.architecture == "modular" {
		return filepath.Join(r.projectRoot, "internal", "platform", "server", "server.go")
	}
	return filepath.Join(r.projectRoot, "internal", "server", "server.go")
}

func (r *RouterAddon) generateServerContent() string {
	switch r.routerType {
	case "gin":
		return r.generateGinServer()
	case "chi":
		return r.generateChiServer()
	case "echo":
		return r.generateEchoServer()
	default:
		return ""
	}
}

func (r *RouterAddon) generateGinServer() string {
	return `package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Server representa el servidor HTTP
type Server struct {
	config *Config
	router *gin.Engine
	server *http.Server
}

// Config contiene la configuración del servidor
type Config struct {
	Port string
}

// NewServer crea una nueva instancia del servidor
func NewServer(config *Config) *Server {
	// Modo release en producción
	// gin.SetMode(gin.ReleaseMode)
	
	router := gin.Default()
	
	return &Server{
		config: config,
		router: router,
		server: &http.Server{
			Addr:    ":" + config.Port,
			Handler: router,
		},
	}
}

// Start inicia el servidor
func (s *Server) Start() error {
	log.Printf("🚀 Servidor iniciado en http://localhost:%s", s.config.Port)
	return s.server.ListenAndServe()
}

// Shutdown detiene el servidor gracefully
func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Deteniendo servidor...")
	return s.server.Shutdown(ctx)
}

// Router retorna el router para configurar rutas
func (s *Server) Router() *gin.Engine {
	return s.router
}
`
}

func (r *RouterAddon) generateChiServer() string {
	return `package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Server representa el servidor HTTP
type Server struct {
	config *Config
	router *chi.Mux
	server *http.Server
}

// Config contiene la configuración del servidor
type Config struct {
	Port string
}

// NewServer crea una nueva instancia del servidor
func NewServer(config *Config) *Server {
	router := chi.NewRouter()
	
	// Middleware básico
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	
	return &Server{
		config: config,
		router: router,
		server: &http.Server{
			Addr:    ":" + config.Port,
			Handler: router,
		},
	}
}

// Start inicia el servidor
func (s *Server) Start() error {
	log.Printf("🚀 Servidor iniciado en http://localhost:%s", s.config.Port)
	return s.server.ListenAndServe()
}

// Shutdown detiene el servidor gracefully
func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Deteniendo servidor...")
	return s.server.Shutdown(ctx)
}

// Router retorna el router para configurar rutas
func (s *Server) Router() *chi.Mux {
	return s.router
}
`
}

func (r *RouterAddon) generateEchoServer() string {
	return `package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server representa el servidor HTTP
type Server struct {
	config *Config
	echo   *echo.Echo
}

// Config contiene la configuración del servidor
type Config struct {
	Port string
}

// NewServer crea una nueva instancia del servidor
func NewServer(config *Config) *Server {
	e := echo.New()
	
	// Middleware básico
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	// Ocultar banner
	e.HideBanner = true
	
	return &Server{
		config: config,
		echo:   e,
	}
}

// Start inicia el servidor
func (s *Server) Start() error {
	log.Printf("🚀 Servidor iniciado en http://localhost:%s", s.config.Port)
	return s.echo.Start(":" + s.config.Port)
}

// Shutdown detiene el servidor gracefully
func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Deteniendo servidor...")
	return s.echo.Shutdown(ctx)
}

// Echo retorna la instancia de Echo para configurar rutas
func (s *Server) Echo() *echo.Echo {
	return s.echo
}
`
}
