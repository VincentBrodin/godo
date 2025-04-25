package utils

import (
	"fmt"
	"os"
)

// Tries to find a file based on name and read it.
// (In short we dont care about extension)
// Name: The name of the file without any path or extension
// Exclude: Any extension that should be ignored, for example .exe or .o
func ReadByName(name string, exclude ...string) ([]byte, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(wd)
	if err != nil {
		return nil, err
	}

outer:
	for _, entrie := range entries {
		if !entrie.Type().IsRegular() || CutExtension(entrie.Name()) != name {
			continue
		}

		extension := GetExtension(entrie.Name())

		for _, ex := range exclude {
			if ex == extension {
				continue outer
			}
		}

		info, err := entrie.Info()
		if err != nil {
			return nil, err
		}

		return os.ReadFile(info.Name()) 
	}

	return nil, fmt.Errorf("Could not find any file with the name: %s\n", name)
}

func GetExtension(name string) string {
	namec := len(name)
	for _i := range namec {
		i := namec - _i - 1
		if name[i] == '.' {
			return name[i:]
		}
	}
	return name
}

func CutExtension(name string) string {
	namec := len(name)
	for _i := range namec {
		i := namec - _i - 1
		if name[i] == '.' {
			return name[:i]
		}
	}
	return name
}
