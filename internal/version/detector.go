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
	return Version{Major: 0, Minor: 1, Patch: 0}, nil
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
	changes := []string{}

	// Definir cambios por versión
	versionChanges := map[string][]string{
		"0.2.0": {
			"✨ Añadido soporte para helpers (response, validator, logger, errors, context)",
			"📦 Helpers integrados en internal/helpers/",
		},
		"0.3.0": {
			"🏗️ Mejoras en la estructura del proyecto",
			"📚 Documentación extendida",
		},
		"0.4.0": {
			"🎯 Comando 'loom generate' para crear componentes individuales",
			"📦 Generación de módulos, handlers, services, models y middlewares",
			"🔍 Detección automática de arquitectura",
		},
		"0.5.0": {
			"⬆️ Comando 'loom upgrade' para actualizar proyectos",
			"💾 Sistema de backup automático antes de actualizar",
			"📊 Detección de versión del proyecto",
		},
	}

	// Recopilar cambios en el rango
	for major := from.Major; major <= to.Major; major++ {
		startMinor := 0
		if major == from.Major {
			startMinor = from.Minor + 1
		}

		endMinor := 99
		if major == to.Major {
			endMinor = to.Minor
		}

		for minor := startMinor; minor <= endMinor; minor++ {
			key := fmt.Sprintf("%d.%d.0", major, minor)
			if changeList, ok := versionChanges[key]; ok {
				changes = append(changes, fmt.Sprintf("\n📌 Versión %s:", key))
				changes = append(changes, changeList...)
			}
		}
	}

	if len(changes) == 0 {
		return "No hay cambios registrados entre estas versiones."
	}

	return strings.Join(changes, "\n")
}
