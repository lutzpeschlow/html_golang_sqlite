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

// ======================================================================================
//
// # GetEnabled
//
// check out main enablement
// READ or WRITE
//
// ======================================================================================
func GetEnabled(obj *objects.Control) string {
	if obj.Enable.WRITE {
		return "WRITE"
	}
	if obj.Enable.READ {
		return "READ"
	}
	return "NONE"
}

// ======================================================================================
//
// # GetReadOption
//
// check out secondary enablement for READ
// - COMMON
// - NUM_ENTRIES
//
// ======================================================================================

func GetReadOption(obj *objects.Control) string {
	if obj.Enable.READOptions.COMMON {
		return "COMMON"
	}
	if obj.Enable.READOptions.NumEntries {
		return "NUM_ENTRIES"
	}
	return "NONE"
}
