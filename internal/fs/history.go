package fs

import (
	"os"
	"path/filepath"
	"sync"
)

type History struct {
	dirs []string
	mu   sync.RWMutex
}

func NewHistory(currentPath string) *History {
	h := &History{
		dirs: make([]string, 0, 32),
	}
	h.initializeFromPath(currentPath)
	return h
}

func (h *History) initializeFromPath(currentPath string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	cleanPath := filepath.Clean(currentPath)
	if cleanPath == "." {
		var err error
		cleanPath, err = os.Getwd()
		if err != nil {
			return
		}
	}

	absPath, err := filepath.Abs(cleanPath)
	if err != nil {
		return
	}

	var pathComponents []string
	current := filepath.Dir(absPath)

	for current != "/" && current != "." {
		pathComponents = append([]string{current}, pathComponents...)
		parent := filepath.Dir(current)

		if parent == current {
			break
		}
		current = parent
	}

	if absPath != "/" && !filepath.IsAbs(absPath) || len(pathComponents) > 0 {
		pathComponents = append([]string{"/"}, pathComponents...)
	}

	h.dirs = pathComponents
}

func (h *History) Push(dir string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if len(h.dirs) > 0 && h.dirs[len(h.dirs)-1] == dir {
		return
	}

	h.dirs = append(h.dirs, dir)
}

func (h *History) Pop() (string, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if len(h.dirs) == 0 {
		return "", false
	}

	dir := h.dirs[len(h.dirs)-1]
	h.dirs = h.dirs[:len(h.dirs)-1]
	return dir, true
}

func (h *History) Peek() (string, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.dirs) == 0 {
		return "", false
	}

	return h.dirs[len(h.dirs)-1], true
}

func (h *History) HasPrevious() bool {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return len(h.dirs) > 0
}

func (h *History) Clear() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.dirs = h.dirs[:0]
}
