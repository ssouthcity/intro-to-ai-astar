package maps

import (
	"io"
	"os"
)

// Interface for any map that can be loaded from an external source
// The intent was to have maps with boolean cells and int cells implement this,
// but a single map struct was already enough to solve all 5 tasks
type MapLoadable interface {
	Load(in io.Reader) error
}

// Helper function for loading a map from a file source
func LoadFromFile(loadable MapLoadable, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = loadable.Load(f)
	if err != nil {
		return err
	}

	return nil
}
