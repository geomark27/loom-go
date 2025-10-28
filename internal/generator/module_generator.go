package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ModuleGenerator genera módulos completos o componentes individuales
type ModuleGenerator struct {
	project *ProjectInfo
}

// NewModuleGenerator crea una nueva instancia del generador de módulos
func NewModuleGenerator(project *ProjectInfo) *ModuleGenerator {
	return &ModuleGenerator{
		project: project,
	}
}

// GenerateModule genera un módulo completo
func (g *ModuleGenerator) GenerateModule(name string, force bool, dryRun bool) ([]string, error) {
	var files []string

	if g.project.Architecture == "layered" {
		files = g.generateLayeredModule(name, force, dryRun)
	} else {
		files = g.generateModularModule(name, force, dryRun)
	}

	return files, nil
}

// generateLayeredModule genera módulo en arquitectura por capas
func (g *ModuleGenerator) generateLayeredModule(name string, force bool, dryRun bool) []string {
	files := []string{}

	nameLower := strings.ToLower(name)
	nameTitle := strings.Title(nameLower)

	// Definir archivos a crear
	filesToCreate := map[string]string{
		fmt.Sprintf("internal/app/handlers/%s_handler.go", nameLower):        g.getHandlerTemplate(nameTitle, nameLower),
		fmt.Sprintf("internal/app/services/%s_service.go", nameLower):        g.getServiceTemplate(nameTitle, nameLower),
		fmt.Sprintf("internal/app/repositories/%s_repository.go", nameLower): g.getRepositoryTemplate(nameTitle, nameLower),
		fmt.Sprintf("internal/app/models/%s.go", nameLower):                  g.getModelTemplate(nameTitle, nameLower),
		fmt.Sprintf("internal/app/dtos/%s_dto.go", nameLower):                g.getDTOTemplate(nameTitle, nameLower),
	}

	// Generar cada archivo
	for filePath, content := range filesToCreate {
		if err := g.createFile(filePath, content, force, dryRun); err != nil {
			fmt.Printf("⚠️  %s: %v\n", filePath, err)
			continue
		}
		files = append(files, filePath)
	}

	return files
}

// generateModularModule genera módulo en arquitectura modular
func (g *ModuleGenerator) generateModularModule(name string, force bool, dryRun bool) []string {
	files := []string{}

	nameLower := strings.ToLower(name)
	nameTitle := strings.Title(nameLower)
	moduleDir := filepath.Join("internal/modules", nameLower)

	// Crear directorio del módulo
	if !dryRun {
		os.MkdirAll(moduleDir, 0755)
	}

	// Definir archivos a crear
	filesToCreate := map[string]string{
		filepath.Join(moduleDir, "handler.go"):    g.getModularHandlerTemplate(nameTitle, nameLower),
		filepath.Join(moduleDir, "service.go"):    g.getModularServiceTemplate(nameTitle, nameLower),
		filepath.Join(moduleDir, "repository.go"): g.getModularRepositoryTemplate(nameTitle, nameLower),
		filepath.Join(moduleDir, "model.go"):      g.getModularModelTemplate(nameTitle, nameLower),
		filepath.Join(moduleDir, "dto.go"):        g.getModularDTOTemplate(nameTitle, nameLower),
		filepath.Join(moduleDir, "module.go"):     g.getModularModuleTemplate(nameTitle, nameLower),
		filepath.Join(moduleDir, "ports.go"):      g.getModularPortsTemplate(nameTitle, nameLower),
		filepath.Join(moduleDir, "errors.go"):     g.getModularErrorsTemplate(nameTitle, nameLower),
	}

	// Generar cada archivo
	for filePath, content := range filesToCreate {
		if err := g.createFile(filePath, content, force, dryRun); err != nil {
			fmt.Printf("⚠️  %s: %v\n", filePath, err)
			continue
		}
		files = append(files, filePath)
	}

	return files
}

// createFile crea un archivo con el contenido dado
func (g *ModuleGenerator) createFile(filePath, content string, force bool, dryRun bool) error {
	// Verificar si el archivo ya existe
	if _, err := os.Stat(filePath); err == nil && !force {
		return fmt.Errorf("ya existe (usa --force para sobrescribir)")
	}

	if dryRun {
		return nil
	}

	// Crear directorio si no existe
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Escribir archivo
	return os.WriteFile(filePath, []byte(content), 0644)
}

// Templates para arquitectura Layered

func (g *ModuleGenerator) getHandlerTemplate(nameTitle, nameLower string) string {
	return fmt.Sprintf(`package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"%s/internal/app/dtos"
	"%s/internal/app/services"
)

type %sHandler struct {
	service *services.%sService
}

func New%sHandler(service *services.%sService) *%sHandler {
	return &%sHandler{
		service: service,
	}
}

// List obtiene todos los %s
func (h *%sHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// GetByID obtiene un %s por ID
func (h *%sHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	item, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// Create crea un nuevo %s
func (h *%sHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dtos.Create%sDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item, err := h.service.Create(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// Update actualiza un %s
func (h *%sHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var dto dtos.Update%sDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item, err := h.service.Update(id, &dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// Delete elimina un %s
func (h *%sHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
`, g.project.ModuleName, g.project.ModuleName, nameTitle, nameTitle, nameTitle, nameTitle, nameTitle, nameTitle,
		nameLower, nameTitle, nameLower, nameTitle, nameLower, nameTitle, nameTitle,
		nameLower, nameTitle, nameTitle, nameLower, nameTitle)
}

func (g *ModuleGenerator) getServiceTemplate(nameTitle, nameLower string) string {
	return fmt.Sprintf(`package services

import (
	"fmt"

	"%s/internal/app/dtos"
	"%s/internal/app/models"
	"%s/internal/app/repositories"
)

type %sService struct {
	repo *repositories.%sRepository
}

func New%sService(repo *repositories.%sRepository) *%sService {
	return &%sService{
		repo: repo,
	}
}

func (s *%sService) GetAll() ([]*models.%s, error) {
	return s.repo.FindAll()
}

func (s *%sService) GetByID(id int) (*models.%s, error) {
	item, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("%s not found")
	}
	return item, nil
}

func (s *%sService) Create(dto *dtos.Create%sDTO) (*models.%s, error) {
	item := &models.%s{
		Name: dto.Name,
		// TODO: Mapear más campos según tu DTO
	}

	return s.repo.Create(item)
}

func (s *%sService) Update(id int, dto *dtos.Update%sDTO) (*models.%s, error) {
	item, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("%s not found")
	}

	if dto.Name != nil {
		item.Name = *dto.Name
	}
	// TODO: Actualizar más campos según tu DTO

	return s.repo.Update(item)
}

func (s *%sService) Delete(id int) error {
	return s.repo.Delete(id)
}
`, g.project.ModuleName, g.project.ModuleName, g.project.ModuleName, nameTitle, nameTitle, nameTitle, nameTitle, nameTitle, nameTitle,
		nameTitle, nameTitle, nameTitle, nameTitle, nameLower, nameTitle, nameTitle, nameTitle, nameTitle,
		nameTitle, nameTitle, nameTitle, nameLower, nameTitle)
}

func (g *ModuleGenerator) getRepositoryTemplate(nameTitle, nameLower string) string {
	return fmt.Sprintf(`package repositories

import (
	"fmt"
	"sync"

	"%s/internal/app/models"
)

type %sRepository struct {
	data   map[int]*models.%s
	nextID int
	mu     sync.RWMutex
}

func New%sRepository() *%sRepository {
	return &%sRepository{
		data:   make(map[int]*models.%s),
		nextID: 1,
	}
}

func (r *%sRepository) FindAll() ([]*models.%s, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	items := make([]*models.%s, 0, len(r.data))
	for _, item := range r.data {
		items = append(items, item)
	}

	return items, nil
}

func (r *%sRepository) FindByID(id int) (*models.%s, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item, exists := r.data[id]
	if !exists {
		return nil, fmt.Errorf("%s not found")
	}

	return item, nil
}

func (r *%sRepository) Create(item *models.%s) (*models.%s, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	item.ID = r.nextID
	r.nextID++

	r.data[item.ID] = item

	return item, nil
}

func (r *%sRepository) Update(item *models.%s) (*models.%s, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[item.ID]; !exists {
		return nil, fmt.Errorf("%s not found")
	}

	r.data[item.ID] = item

	return item, nil
}

func (r *%sRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[id]; !exists {
		return fmt.Errorf("%s not found")
	}

	delete(r.data, id)

	return nil
}
`, g.project.ModuleName, nameTitle, nameTitle, nameTitle, nameTitle, nameTitle, nameTitle,
		nameTitle, nameTitle, nameTitle, nameTitle, nameTitle, nameLower,
		nameTitle, nameTitle, nameTitle, nameTitle, nameTitle, nameTitle, nameLower,
		nameTitle, nameLower)
}

func (g *ModuleGenerator) getModelTemplate(nameTitle, nameLower string) string {
	return fmt.Sprintf(`package models

import "time"

type %s struct {
	ID        int       `+"`json:\"id\"`"+`
	Name      string    `+"`json:\"name\"`"+`
	CreatedAt time.Time `+"`json:\"created_at\"`"+`
	UpdatedAt time.Time `+"`json:\"updated_at\"`"+`
	// TODO: Agregar más campos según tus necesidades
}
`, nameTitle)
}

func (g *ModuleGenerator) getDTOTemplate(nameTitle, nameLower string) string {
	return fmt.Sprintf(`package dtos

type Create%sDTO struct {
	Name string `+"`json:\"name\" binding:\"required\"`"+`
	// TODO: Agregar más campos según tus necesidades
}

type Update%sDTO struct {
	Name *string `+"`json:\"name,omitempty\"`"+`
	// TODO: Agregar más campos según tus necesidades
}
`, nameTitle, nameTitle)
}

// Templates para arquitectura Modular

func (g *ModuleGenerator) getModularHandlerTemplate(nameTitle, nameLower string) string {
	return fmt.Sprintf(`package %s

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// RegisterRoutes registra las rutas del módulo
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/%s", h.List).Methods("GET")
	router.HandleFunc("/%s/{id}", h.GetByID).Methods("GET")
	router.HandleFunc("/%s", h.Create).Methods("POST")
	router.HandleFunc("/%s/{id}", h.Update).Methods("PUT")
	router.HandleFunc("/%s/{id}", h.Delete).Methods("DELETE")
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	item, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var dto Create%sDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item, err := h.service.Create(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var dto Update%sDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item, err := h.service.Update(id, &dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
`, nameLower, nameLower, nameLower, nameLower, nameLower, nameLower, nameTitle, nameTitle)
}

func (g *ModuleGenerator) getModularServiceTemplate(nameTitle, nameLower string) string {
	return fmt.Sprintf(`package %s

type ServiceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &ServiceImpl{
		repo: repo,
	}
}

func (s *ServiceImpl) GetAll() ([]*%s, error) {
	return s.repo.FindAll()
}

func (s *ServiceImpl) GetByID(id int) (*%s, error) {
	return s.repo.FindByID(id)
}

func (s *ServiceImpl) Create(dto *Create%sDTO) (*%s, error) {
	item := &%s{
		Name: dto.Name,
		// TODO: Mapear más campos
	}

	return s.repo.Create(item)
}

func (s *ServiceImpl) Update(id int, dto *Update%sDTO) (*%s, error) {
	item, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if dto.Name != nil {
		item.Name = *dto.Name
	}
	// TODO: Actualizar más campos

	return s.repo.Update(item)
}

func (s *ServiceImpl) Delete(id int) error {
	return s.repo.Delete(id)
}
`, nameLower, nameTitle, nameTitle, nameTitle, nameTitle, nameTitle, nameTitle, nameTitle)
}

func (g *ModuleGenerator) getModularRepositoryTemplate(nameTitle, nameLower string) string {
	return fmt.Sprintf(`package %s

import (
	"fmt"
	"sync"
)

type RepositoryImpl struct {
	data   map[int]*%s
	nextID int
	mu     sync.RWMutex
}

func NewRepository() Repository {
	return &RepositoryImpl{
		data:   make(map[int]*%s),
		nextID: 1,
	}
}

func (r *RepositoryImpl) FindAll() ([]*%s, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	items := make([]*%s, 0, len(r.data))
	for _, item := range r.data {
		items = append(items, item)
	}

	return items, nil
}

func (r *RepositoryImpl) FindByID(id int) (*%s, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item, exists := r.data[id]
	if !exists {
		return nil, ErrNotFound
	}

	return item, nil
}

func (r *RepositoryImpl) Create(item *%s) (*%s, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	item.ID = r.nextID
	r.nextID++

	r.data[item.ID] = item

	return item, nil
}

func (r *RepositoryImpl) Update(item *%s) (*%s, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[item.ID]; !exists {
		return nil, ErrNotFound
	}

	r.data[item.ID] = item

	return item, nil
}

func (r *RepositoryImpl) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[id]; !exists {
		return ErrNotFound
	}

	delete(r.data, id)

	return nil
}
`, nameLower, nameTitle, nameTitle, nameTitle, nameTitle, nameTitle, nameTitle, nameTitle, nameTitle, nameTitle)
}

func (g *ModuleGenerator) getModularModelTemplate(nameTitle, nameLower string) string {
	return fmt.Sprintf(`package %s

import "time"

type %s struct {
	ID        int       `+"`json:\"id\"`"+`
	Name      string    `+"`json:\"name\"`"+`
	CreatedAt time.Time `+"`json:\"created_at\"`"+`
	UpdatedAt time.Time `+"`json:\"updated_at\"`"+`
	// TODO: Agregar más campos
}
`, nameLower, nameTitle)
}

func (g *ModuleGenerator) getModularDTOTemplate(nameTitle, nameLower string) string {
	return fmt.Sprintf(`package %s

type Create%sDTO struct {
	Name string `+"`json:\"name\" binding:\"required\"`"+`
	// TODO: Agregar más campos
}

type Update%sDTO struct {
	Name *string `+"`json:\"name,omitempty\"`"+`
	// TODO: Agregar más campos
}
`, nameLower, nameTitle, nameTitle)
}

func (g *ModuleGenerator) getModularModuleTemplate(nameTitle, nameLower string) string {
	return fmt.Sprintf(`package %s

import "github.com/gorilla/mux"

type Module struct {
	handler *Handler
}

func NewModule() *Module {
	repo := NewRepository()
	service := NewService(repo)
	handler := NewHandler(service)

	return &Module{
		handler: handler,
	}
}

func (m *Module) RegisterRoutes(router *mux.Router) {
	m.handler.RegisterRoutes(router.PathPrefix("/api/v1").Subrouter())
}
`, nameLower)
}

func (g *ModuleGenerator) getModularPortsTemplate(nameTitle, nameLower string) string {
	return fmt.Sprintf(`package %s

// Service define los métodos de negocio del módulo
type Service interface {
	GetAll() ([]*%s, error)
	GetByID(id int) (*%s, error)
	Create(dto *Create%sDTO) (*%s, error)
	Update(id int, dto *Update%sDTO) (*%s, error)
	Delete(id int) error
}

// Repository define los métodos de persistencia del módulo
type Repository interface {
	FindAll() ([]*%s, error)
	FindByID(id int) (*%s, error)
	Create(item *%s) (*%s, error)
	Update(item *%s) (*%s, error)
	Delete(id int) error
}
`, nameLower, nameTitle, nameTitle, nameTitle, nameTitle, nameTitle, nameTitle,
		nameTitle, nameTitle, nameTitle, nameTitle, nameTitle, nameTitle)
}

func (g *ModuleGenerator) getModularErrorsTemplate(nameTitle, nameLower string) string {
	return fmt.Sprintf(`package %s

import "errors"

var (
	ErrNotFound      = errors.New("%s not found")
	ErrInvalidInput  = errors.New("invalid input")
	ErrAlreadyExists = errors.New("%s already exists")
)
`, nameLower, nameLower, nameLower)
}

// GenerateHandler genera solo el archivo handler
func (g *ModuleGenerator) GenerateHandler(name string, force bool, dryRun bool) ([]string, error) {
	nameLower := strings.ToLower(name)
	nameTitle := strings.Title(nameLower)

	var filePath string
	var content string

	if g.project.Architecture == "layered" {
		filePath = fmt.Sprintf("internal/app/handlers/%s_handler.go", nameLower)
		content = g.getHandlerTemplate(nameTitle, nameLower)
	} else {
		filePath = fmt.Sprintf("internal/modules/%s/handler.go", nameLower)
		content = g.getModularHandlerTemplate(nameTitle, nameLower)
	}

	if err := g.createFile(filePath, content, force, dryRun); err != nil {
		return nil, err
	}

	return []string{filePath}, nil
}

// GenerateService genera solo el archivo service
func (g *ModuleGenerator) GenerateService(name string, force bool, dryRun bool) ([]string, error) {
	nameLower := strings.ToLower(name)
	nameTitle := strings.Title(nameLower)

	var filePath string
	var content string

	if g.project.Architecture == "layered" {
		filePath = fmt.Sprintf("internal/app/services/%s_service.go", nameLower)
		content = g.getServiceTemplate(nameTitle, nameLower)
	} else {
		filePath = fmt.Sprintf("internal/modules/%s/service.go", nameLower)
		content = g.getModularServiceTemplate(nameTitle, nameLower)
	}

	if err := g.createFile(filePath, content, force, dryRun); err != nil {
		return nil, err
	}

	return []string{filePath}, nil
}

// GenerateModel genera solo el archivo model
func (g *ModuleGenerator) GenerateModel(name string, force bool, dryRun bool) ([]string, error) {
	nameLower := strings.ToLower(name)
	nameTitle := strings.Title(nameLower)

	var filePath string
	var content string

	if g.project.Architecture == "layered" {
		filePath = fmt.Sprintf("internal/app/models/%s.go", nameLower)
		content = g.getModelTemplate(nameTitle, nameLower)
	} else {
		filePath = fmt.Sprintf("internal/modules/%s/model.go", nameLower)
		content = g.getModularModelTemplate(nameTitle, nameLower)
	}

	if err := g.createFile(filePath, content, force, dryRun); err != nil {
		return nil, err
	}

	return []string{filePath}, nil
}

// GenerateMiddleware genera un middleware
func (g *ModuleGenerator) GenerateMiddleware(name string, force bool, dryRun bool) ([]string, error) {
	nameLower := strings.ToLower(name)
	nameTitle := strings.Title(nameLower)

	var filePath string
	var content string

	if g.project.Architecture == "layered" {
		filePath = fmt.Sprintf("internal/app/middleware/%s.go", nameLower)
	} else {
		filePath = fmt.Sprintf("internal/middleware/%s.go", nameLower)
	}

	content = fmt.Sprintf(`package middleware

import (
	"log"
	"net/http"
)

// %s middleware
func %s(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Middleware %s: %%s %%s", r.Method, r.URL.Path)
		
		// Implementa tu lógica aquí
		
		next.ServeHTTP(w, r)
	})
}
`, nameTitle, nameTitle, nameLower)

	if err := g.createFile(filePath, content, force, dryRun); err != nil {
		return nil, err
	}

	return []string{filePath}, nil
}
