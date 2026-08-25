package site

import (
	"embed"
	"net/http"
)

//go:embed index.html favicon.ico img/* images/* fonts/* css/* js/*
var Assets embed.FS

type tryFileHandler struct {
	root        http.FileSystem
	defaultFile string // 回退文件，一般 "index.html"
}

// TryFileServer 构造 try_files handler
// root：静态文件源，embed.FS 需要转 http.FileSystem
// index：回退的默认文件
func NewTryFileServer(root http.FileSystem, index string) *tryFileHandler {
	return &tryFileHandler{
		root:        root,
		defaultFile: index,
	}
}

func (h *tryFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.Path

	// 1.尝试打开请求的资源
	f, err := h.root.Open(requestPath)
	if err == nil {
		defer f.Close()
		fi, errStat := f.Stat()
		if errStat == nil {
			// 文件存在，且不是目录 → 直接提供该文件
			if !fi.IsDir() {
				http.ServeContent(w, r, requestPath, fi.ModTime(), f)
				return
			}
		}
	}

	// 走到这里：文件不存在 OR 是目录 → 回退 defaultFile(index.html)
	defFile, errDef := h.root.Open(h.defaultFile)
	if errDef != nil {
		http.Error(w, "not found, default file missing", http.StatusNotFound)
		return
	}
	defer defFile.Close()
	defFi, _ := defFile.Stat()
	http.ServeContent(w, r, h.defaultFile, defFi.ModTime(), defFile)
}
