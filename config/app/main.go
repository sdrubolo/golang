package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	gocache "github.com/patrickmn/go-cache"
)

type Songs struct {
	Songs []Album `json:"songs"`
}

type Album struct {
	Title     string `json:"title"`
	Artist    string `json:"artist"`
	Year      string `json:"year"`
	WebURL    string `json:"web_url"`
	Thumbnail string `json:"img_url"`
}

type FileWatcher struct {
	store *gocache.Cache
	dir   string
	file  string
}

func (f *FileWatcher) readSongs() error {
	fmt.Printf("readSongs called for %s\n", filepath.Base(f.dir+"/"+f.file))
	file := f.dir + "/" + f.file
	content, err := os.ReadFile(file)
	// Convert the byte slice to a string and print it

	if err != nil {
		return err
	}

	var songs Songs
	err = json.Unmarshal(content, &songs)

	if err != nil {
		fmt.Printf("readSongs error during Unmarshal %s\n", err)
		return err
	}
	f.store.Set("file-content", songs, -1)

	return nil
}

func (f *FileWatcher) watch() (*fsnotify.Watcher, error) {

	var (
		file = filepath.Clean(f.dir + "/" + f.file)
		// Wait 100ms for new events; each new event resets the timer.
		waitFor = 100 * time.Millisecond

		// Keep track of the timers, as path → timer.
		mu     sync.Mutex
		timers = make(map[string]*time.Timer)

		// Callback we run.
		handleEvent = func(e fsnotify.Event) {
			err := f.readSongs()
			if err != nil {
				return
			}

			// Don't need to remove the timer if you don't have a lot of files.
			mu.Lock()
			delete(timers, e.Name)
			mu.Unlock()
		}
	)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	err = watcher.Add(f.dir)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Printf("event received for file %s => %s\n", file, event)
					if event.Name != file {
						continue
					}
					// Get timer.
					mu.Lock()
					t, ok := timers[event.Name]
					mu.Unlock()

					// No timer yet, so create one.
					if !ok {
						t = time.AfterFunc(math.MaxInt64, func() { handleEvent(event) })
						t.Stop()

						mu.Lock()
						timers[event.Name] = t
						mu.Unlock()
					}
					t.Reset(waitFor)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("Error:", err)
			}
		}
	}()

	return watcher, nil
}

func main() {
	filepath := "./songs"
	filename := "albums.json"

	goCache := gocache.New(gocache.NoExpiration, 5*time.Minute)

	fileWatcher := FileWatcher{
		store: goCache,
		dir:   filepath,
		file:  filename,
	}

	err := fileWatcher.readSongs()
	if err != nil {
		log.Fatal(err)
	}
	watcher, err := fileWatcher.watch()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	apiHandler := ApiHandler{
		store: goCache,
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/list", apiHandler.list)

	log.Fatal(http.ListenAndServe(":8080", r))
}
