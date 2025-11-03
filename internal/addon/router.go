package addon

import (
	"fmt"
	"path/filepath"
	"strings"
)

// RouterAddon manages HTTP router installation
type RouterAddon struct {
	projectRoot  string
	architecture string
	routerType   string // "gin", "chi", "echo"
}

// NewRouterAddon creates a new router addon
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
		"gin":  "Fast web framework with excellent performance",
		"chi":  "Lightweight and composable router compatible with net/http",
		"echo": "Minimalist high-performance framework",
	}
	return descriptions[r.routerType]
}

func (r *RouterAddon) IsInstalled() (bool, error) {
	detector := NewProjectDetector(r.projectRoot)
	currentRouter := detector.DetectRouter()
	return currentRouter == r.routerType, nil
}

func (r *RouterAddon) CanInstall() (bool, string, error) {
	// Routers can always be installed (they replace the existing one)
	return true, "", nil
}

func (r *RouterAddon) GetConflicts() []string {
	// All routers conflict with each other
	conflicts := []string{"gin", "chi", "echo", "gorilla-mux"}

	// Remove the current router from conflicts
	filtered := []string{}
	for _, c := range conflicts {
		if c != r.routerType {
			filtered = append(filtered, c)
		}
	}

	return filtered
}

func (r *RouterAddon) Install(force bool) error {
	// 1. Update go.mod with the dependency
	if err := r.addDependency(); err != nil {
		return fmt.Errorf("error adding dependency: %w", err)
	}

	// 2. Update server files
	if err := r.updateServerFiles(); err != nil {
		return fmt.Errorf("error updating server: %w", err)
	}

	// 3. Update handlers (if necessary)
	fmt.Println("⚠️  Note: Existing handlers will need to be manually updated")
	fmt.Println("   to use the new router API")

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
		return fmt.Errorf("invalid module format")
	}

	fmt.Printf("   📦 Adding dependency: %s\n", modules[r.routerType])
	return UpdateGoMod(parts[0], parts[1])
}

func (r *RouterAddon) updateServerFiles() error {
	serverPath := r.getServerPath()

	fmt.Printf("   📝 Updating %s\n", serverPath)

	// Generate new content according to the router
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

// Server represents the HTTP server
type Server struct {
	config *Config
	router *gin.Engine
	server *http.Server
}

// Config contains the server configuration
type Config struct {
	Port string
}

// NewServer creates a new server instance
func NewServer(config *Config) *Server {
	// Release mode in production
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

// Start starts the server
func (s *Server) Start() error {
	log.Printf("🚀 Server started at http://localhost:%s", s.config.Port)
	return s.server.ListenAndServe()
}

// Shutdown stops the server gracefully
func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Stopping server...")
	return s.server.Shutdown(ctx)
}

// Router returns the router to configure routes
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

// Server represents the HTTP server
type Server struct {
	config *Config
	router *chi.Mux
	server *http.Server
}

// Config contains the server configuration
type Config struct {
	Port string
}

// NewServer creates a new server instance
func NewServer(config *Config) *Server {
	router := chi.NewRouter()

	// Basic middleware
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

// Start starts the server
func (s *Server) Start() error {
	log.Printf("🚀 Server started at http://localhost:%s", s.config.Port)
	return s.server.ListenAndServe()
}

// Shutdown stops the server gracefully
func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Stopping server...")
	return s.server.Shutdown(ctx)
}

// Router returns the router to configure routes
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

// Server represents the HTTP server
type Server struct {
	config *Config
	echo   *echo.Echo
}

// Config contains the server configuration
type Config struct {
	Port string
}

// NewServer creates a new server instance
func NewServer(config *Config) *Server {
	e := echo.New()

	// Basic middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Hide banner
	e.HideBanner = true

	return &Server{
		config: config,
		echo:   e,
	}
}

// Start starts the server
func (s *Server) Start() error {
	log.Printf("🚀 Server started at http://localhost:%s", s.config.Port)
	return s.echo.Start(":" + s.config.Port)
}

// Shutdown stops the server gracefully
func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Stopping server...")
	return s.echo.Shutdown(ctx)
}

// Echo returns the Echo instance to configure routes
func (s *Server) Echo() *echo.Echo {
	return s.echo
}
`
}
