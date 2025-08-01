package helpers

import (
	"fmt"
	"path/filepath"
)

func Include(path string) []string {
	files, err := filepath.Glob("admin/views/templates/*html")
	if err != nil {
		fmt.Println(err)
	}
	path_files, err := filepath.Glob("admin/views/" + path + "/*.html")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range path_files {
		files = append(files, file)
	}
	return files
}
