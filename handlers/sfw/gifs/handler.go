package sfw_gifs_handler

import (
	"fmt"
	"io/fs"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/catpawzz/fluff-fetcher/utils"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandomGif(category string) (string, error) {
	basePath := "./storage/sfw/gifs"
	folderPath := filepath.Join(basePath, category)
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return "", fmt.Errorf("failed to access category folder: %w", err)
	}
	var gifFiles []fs.DirEntry
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(strings.ToLower(file.Name()), ".gif") {
			gifFiles = append(gifFiles, file)
		}
	}
	if len(gifFiles) == 0 {
		return "", fmt.Errorf("no gif files found in category: %s", category)
	}
	randomGif := gifFiles[rand.Intn(len(gifFiles))]
	return filepath.Join(folderPath, randomGif.Name()), nil
}

func getAvailableCategories() ([]string, error) {
	basePath := "./storage/sfw/gifs"
	entries, err := os.ReadDir(basePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read gifs directory: %w", err)
	}

	var categories []string
	for _, entry := range entries {
		if entry.IsDir() {
			categories = append(categories, entry.Name())
		}
	}
	return categories, nil
}

func SfwGifHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/sfw/gifs/")
	if path == "" {
		categories, err := getAvailableCategories()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"handler": "/handlers/sfw/gifs","title": "Error 500","message": "Failed to list categories.", "status": "failure"}`)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonResponse := fmt.Sprintf(`{"handler": "/handlers/sfw/gifs","title": "Available Categories","categories": [`)
		for i, category := range categories {
			if i > 0 {
				jsonResponse += ", "
			}
			jsonResponse += fmt.Sprintf(`"%s"`, category)
		}
		jsonResponse += `], "status": "success"}`

		fmt.Fprint(w, jsonResponse)
		return
	}
	gifPath, err := getRandomGif(path)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"handler": "/handlers/sfw/gifs","title": "Error 404","message": "Category not found or no GIFs available.", "status": "failure"}`)
		return
	}
	file, err := os.Open(gifPath)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"handler": "/handlers/sfw/gifs","title": "Error 500","message": "Failed to open GIF file.", "status": "failure"}`)
		return
	}
	defer file.Close()
	utils.IncrementServedCounter()

	// Add the file path as a header
	relativePath := strings.TrimPrefix(gifPath, "./storage/")
	w.Header().Set("X-File-Path", relativePath)

	w.Header().Set("Content-Type", "image/gif")
	http.ServeContent(w, r, filepath.Base(gifPath), time.Now(), file)
}
