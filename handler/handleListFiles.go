package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

type M map[string]any

func HandleListFiles(w http.ResponseWriter, r *http.Request) {
	files := []M{}

	// os.Getwd mengembalikan informasi absolute path dimana aplikasi dieksekusi
	basePath, _ := os.Getwd()
	fileLocation := filepath.Join(basePath, "files")

	err := filepath.Walk(fileLocation,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			files = append(files, M{"filename": info.Name(), "path": path})
			return nil
		},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(files)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
