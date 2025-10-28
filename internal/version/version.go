package version

import (
	"fmt"
	"strconv"
	"strings"
)

// Version representa una versión semántica
type Version struct {
	Major int
	Minor int
	Patch int
}

// Current es la versión actual de Loom
var Current = Version{Major: 1, Minor: 0, Patch: 0}

// String retorna la versión en formato string
func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// Parse convierte un string a Version
func Parse(s string) (Version, error) {
	s = strings.TrimPrefix(s, "v")
	parts := strings.Split(s, ".")

	if len(parts) != 3 {
		return Version{}, fmt.Errorf("formato de versión inválido: %s", s)
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return Version{}, fmt.Errorf("major version inválido: %s", parts[0])
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return Version{}, fmt.Errorf("minor version inválido: %s", parts[1])
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return Version{}, fmt.Errorf("patch version inválido: %s", parts[2])
	}

	return Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}, nil
}

// Compare compara dos versiones
// Retorna: -1 si v < other, 0 si v == other, 1 si v > other
func (v Version) Compare(other Version) int {
	if v.Major != other.Major {
		if v.Major < other.Major {
			return -1
		}
		return 1
	}

	if v.Minor != other.Minor {
		if v.Minor < other.Minor {
			return -1
		}
		return 1
	}

	if v.Patch != other.Patch {
		if v.Patch < other.Patch {
			return -1
		}
		return 1
	}

	return 0
}

// IsNewer verifica si esta versión es más nueva que otra
func (v Version) IsNewer(other Version) bool {
	return v.Compare(other) > 0
}

// IsOlder verifica si esta versión es más antigua que otra
func (v Version) IsOlder(other Version) bool {
	return v.Compare(other) < 0
}

// IsCompatible verifica si dos versiones son compatibles (mismo major)
func (v Version) IsCompatible(other Version) bool {
	return v.Major == other.Major
}
