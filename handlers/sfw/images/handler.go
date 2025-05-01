package sfw_images_handler

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

func getRandomImages(category string) (string, error) {
	basePath := "./storage/sfw/images"
	folderPath := filepath.Join(basePath, category)
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return "", fmt.Errorf("failed to access category folder: %w", err)
	}
	var imageFiles []fs.DirEntry
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(strings.ToLower(file.Name()), ".jpg") {
			imageFiles = append(imageFiles, file)
		}
	}
	if len(imageFiles) == 0 {
		return "", fmt.Errorf("no image files found in category: %s", category)
	}
	randomImage := imageFiles[rand.Intn(len(imageFiles))]
	return filepath.Join(folderPath, randomImage.Name()), nil
}

func getAvailableCategories() ([]string, error) {
	basePath := "./storage/sfw/images"
	entries, err := os.ReadDir(basePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read images directory: %w", err)
	}

	var categories []string
	for _, entry := range entries {
		if entry.IsDir() {
			categories = append(categories, entry.Name())
		}
	}
	return categories, nil
}

func SfwImageHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/sfw/images/")
	if path == "" {
		categories, err := getAvailableCategories()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"handler": "/handlers/sfw/images","title": "Error 500","message": "Failed to list categories.", "status": "failure"}`)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonResponse := fmt.Sprintf(`{"handler": "/handlers/sfw/images","title": "Available Categories","categories": [`)
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
	imagePath, err := getRandomImages(path)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"handler": "/handlers/sfw/images","title": "Error 404","message": "Category not found or no images available.", "status": "failure"}`)
		return
	}
	file, err := os.Open(imagePath)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"handler": "/handlers/sfw/images","title": "Error 500","message": "Failed to open image file.", "status": "failure"}`)
		return
	}
	defer file.Close()
	utils.IncrementServedCounter()
	w.Header().Set("Content-Type", "image/jpeg")
	http.ServeContent(w, r, filepath.Base(imagePath), time.Now(), file)
}
