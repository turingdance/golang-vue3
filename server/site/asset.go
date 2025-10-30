package site

import (
	"embed"
	"net/http"
)

// img/* images/* fonts/* css/* js/* tinymce/* *.js *.pdf *.mp4
//
//go:embed index.html favicon.ico
var Assets embed.FS

type tryFileHandler struct {
	root        http.FileSystem
	defaultFile string
}

func TryFileServer(root http.FileSystem, index string) *tryFileHandler {
	return &tryFileHandler{
		root:        root,
		defaultFile: index,
	}
}

// /img/a.js
func (h *tryFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// check whether a file exists or is a directory at the given path
	// h.root.Open(r.RequestURI)
	// fi, err := os.Stat(path)
	// if os.IsNotExist(err) || fi.IsDir() {
	// 	// file does not exist or path is a directory, serve index.html
	// 	http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
	// 	return
	// }

	// if err != nil {
	// 	// if we got an error (that wasn't that the file doesn't exist) stating the
	// 	// file, return a 500 internal server error and stop
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// // otherwise, use http.FileServer to serve the static file
	// http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
