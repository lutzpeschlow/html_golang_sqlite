package ctrl

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/lutzpeschlow/html_golang_sqlite/objects"
)

func ReadControlJsonFile(path string, obj *objects.Control) error {
	// (1) read json control file
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read control file: %w", err)
	}
	// (2) json parsen
	if err := json.Unmarshal(data, obj); err != nil {
		return fmt.Errorf("parse config json %s: %w", path, err)
	}
	// (3) reine Logik in eigene Funktion auslagern
	// return PrepareControl(obj, osName)
	return nil
}

func PrintEnabled(obj *objects.Control) {
	fmt.Println("enabled actions:")
	for name, enabled := range obj.Enable {
		if enabled {
			fmt.Println(" -", name)
		}
	}
}

func GetEnabled(obj *objects.Control) string {
	for name, enabled := range obj.Enable {
		if enabled {
			return name
		}
	}
	return "NONE"
}
