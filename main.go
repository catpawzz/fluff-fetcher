package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"

	sfw_gif_handler "github.com/catpawzz/fluff-fetcher/handlers/sfw/gifs"
	sfw_images_handler "github.com/catpawzz/fluff-fetcher/handlers/sfw/images"
	"github.com/catpawzz/fluff-fetcher/utils"
	"github.com/joho/godotenv"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

var startTime time.Time
var version = "2404252021"
var servedCounter int64

func isBrowserRequest(r *http.Request) bool {
	userAgent := r.Header.Get("User-Agent")
	return strings.Contains(strings.ToLower(userAgent), "mozilla") ||
		strings.Contains(strings.ToLower(userAgent), "safari") ||
		strings.Contains(strings.ToLower(userAgent), "chrome") ||
		strings.Contains(strings.ToLower(userAgent), "edge")
}

func handleBrowserRequest(w http.ResponseWriter, r *http.Request, pagePath string) bool {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	if isBrowserRequest(r) {
		http.ServeFile(w, r, pagePath)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"title": "Error 404","message": "There is no API endpoint at this address, access the api using the /api path!", "status": "failure"}`)
	}
	return false
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	handleBrowserRequest(w, r, "./static/main.html")
}

func cdnHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/cdn/")
	filePath := filepath.Join("./", path)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"handler": "cdn","title": "Error 404","message": "File not found.", "status": "failure"}`)
		return
	}

	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".jpg", ".jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	case ".gif":
		w.Header().Set("Content-Type", "image/gif")
	default:
		w.Header().Set("Content-Type", "application/octet-stream")
	}

	http.ServeFile(w, r, filePath)
	IncrementServedCounter()
}

func main() {
	startTime = time.Now()

	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		uptime := func(d time.Duration) string {
			days := int(d.Hours() / 24)
			if days > 0 {
				return fmt.Sprintf("%d days", days)
			}
			hours := int(d.Hours())
			if hours > 0 {
				return fmt.Sprintf("%d hours", hours)
			}
			minutes := int(d.Minutes())
			if minutes > 0 {
				return fmt.Sprintf("%d minutes", minutes)
			}
			return fmt.Sprintf("%d seconds", int(d.Seconds()))
		}(time.Since(startTime))
		loadStat, err := load.Avg()
		serverLoad := "unknown"
		if err == nil {
			serverLoad = fmt.Sprintf("%.2f, %.2f, %.2f", loadStat.Load1, loadStat.Load5, loadStat.Load15)
		}
		memStat, err := mem.VirtualMemory()
		memoryUsage := "unknown"
		if err == nil {
			memoryUsage = fmt.Sprintf("%.2f%%", memStat.UsedPercent)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"status": "running", 
			"version": "%s", 
			"uptime": "%s", 
			"server_time": "%s", 
			"server_load": "%s",
			"memory_usage": "%s",
			"served_files": %d
		}`, version, uptime, time.Now().Format("2006-01-02 15:04:05"), serverLoad, memoryUsage, utils.GetServedCount())
	})
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/api/sfw/gifs/", func(w http.ResponseWriter, r *http.Request) {
		sfw_gif_handler.SfwGifHandler(w, r)
		IncrementServedCounter()
	})
	http.HandleFunc("/api/sfw/gifs", func(w http.ResponseWriter, r *http.Request) {
		sfw_gif_handler.SfwGifHandler(w, r)
		IncrementServedCounter()
	})
	http.HandleFunc("/api/sfw/images/", func(w http.ResponseWriter, r *http.Request) {
		sfw_images_handler.SfwImageHandler(w, r)
		IncrementServedCounter()
	})
	http.HandleFunc("/api/sfw/images", func(w http.ResponseWriter, r *http.Request) {
		sfw_images_handler.SfwImageHandler(w, r)
		IncrementServedCounter()
	})
	http.HandleFunc("/cdn/", cdnHandler)
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"handler": "main","title": "Error 404","message": "Please specify an existing subpath.", "status": "failure"}`)
	})
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("")
	fmt.Println("      :::::::::: :::       :::    ::: :::::::::: :::::::::: :::::::::: :::::::::: ::::::::::: ::::::::  :::    ::: :::::::::: ::::::::: ")
	fmt.Println("     :+:        :+:       :+:    :+: :+:        :+:        :+:        :+:            :+:    :+:    :+: :+:    :+: :+:        :+:    :+: ")
	fmt.Println("    +:+        +:+       +:+    +:+ +:+        +:+        +:+        +:+            +:+    +:+        +:+    +:+ +:+        +:+    +:+  ")
	fmt.Println("   :#::+::#   +#+       +#+    +:+ :#::+::#   :#::+::#   :#::+::#   +#++:++#       +#+    +#+        +#++:++#++ +#++:++#   +#++:++#:    ")
	fmt.Println("  +#+        +#+       +#+    +#+ +#+        +#+        +#+        +#+            +#+    +#+        +#+    +#+ +#+        +#+    +#+    ")
	fmt.Println(" #+#        #+#       #+#    #+# #+#        #+#        #+#        #+#            #+#    #+#    #+# #+#    #+# #+#        #+#    #+#     ")
	fmt.Println("###        ########## ########  ###        ###        ###        ##########     ###     ########  ###    ### ########## ###    ###      ")
	fmt.Println("")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("This was made by actual catboys :0")
	fmt.Println("Please consult the documentation to run your API correctly")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("EXITING ✘ | Could not load the .env file, does it exist?")
		os.Exit(1)
	}
	fmt.Println("CHECK ✓ | .env file found and loaded")
	fmt.Println("RUNNING ✓ | Fluff-Fetcher server starting on port " + os.Getenv("PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		fmt.Println("EXITING ✘ | Failed to start Fluff-Fetcher server:", err)
		os.Exit(1)
	}
}

func IncrementServedCounter() {
	atomic.AddInt64(&servedCounter, 1)
}
