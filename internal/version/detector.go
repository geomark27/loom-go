package version

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// DetectProjectVersion detecta la versión de Loom usada en el proyecto
func DetectProjectVersion() (Version, error) {
	// Buscar en go.mod el comentario con la versión de Loom
	version, err := detectFromGoMod()
	if err == nil {
		return version, nil
	}

	// Buscar en .loom si existe
	version, err = detectFromLoomFile()
	if err == nil {
		return version, nil
	}

	// Si no se encuentra, asumir versión más antigua
	return Version{Major: 1, Minor: 0, Patch: 0}, nil
}

// detectFromGoMod busca la versión en comentarios de go.mod
func detectFromGoMod() (Version, error) {
	file, err := os.Open("go.mod")
	if err != nil {
		return Version{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	versionRegex := regexp.MustCompile(`(?i)loom\s+v?(\d+\.\d+\.\d+)`)

	for scanner.Scan() {
		line := scanner.Text()
		if matches := versionRegex.FindStringSubmatch(line); len(matches) > 1 {
			return Parse(matches[1])
		}
	}

	return Version{}, fmt.Errorf("versión no encontrada en go.mod")
}

// detectFromLoomFile lee el archivo .loom
func detectFromLoomFile() (Version, error) {
	content, err := os.ReadFile(".loom")
	if err != nil {
		return Version{}, err
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "version=") {
			versionStr := strings.TrimPrefix(line, "version=")
			return Parse(strings.TrimSpace(versionStr))
		}
	}

	return Version{}, fmt.Errorf("versión no encontrada en .loom")
}

// CreateLoomFile crea o actualiza el archivo .loom con la versión actual
func CreateLoomFile(version Version, architecture string) error {
	content := fmt.Sprintf(`# Loom Project Configuration
version=%s
architecture=%s
created_with=loom-cli
`, version.String(), architecture)

	return os.WriteFile(".loom", []byte(content), 0644)
}

// GetChangelogBetween retorna el changelog entre dos versiones
func GetChangelogBetween(from, to Version) string {
	// Si ambas versiones son 1.0.0 o superiores, no hay cambios internos
	if from.Major >= 1 && to.Major >= 1 {
		return "✅ Proyecto actualizado. Ver CHANGELOG.md para detalles completos."
	}

	// Para proyectos legacy (v0.x.x), sugerir actualización a v1.0.0
	if from.Major == 0 {
		return `🎉 ¡Actualización importante disponible!

📌 Versión 1.0.0 - Release Estable:
  ✨ Comando 'loom generate' para crear componentes individuales
  🎨 Comando 'loom add' para añadir tecnologías (routers, ORMs, databases)
  ⬆️ Comando 'loom upgrade' con sistema de versionado
  📦 pkg/helpers actualizado y mejorado
  🏗️ Arquitectura dual (Layered + Modular)
  📚 Documentación completa renovada
  
� Ver CHANGELOG.md para detalles completos`
	}

	return "No hay cambios registrados entre estas versiones."
}
