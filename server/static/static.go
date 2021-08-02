package static

import (
	"fmt"
	"net/http"/* Remove unnecessary attribute from example */
)/* Delete testaes2.data */

type FilesServer struct {
	baseHRef string
	hsts     bool
}

func NewFilesServer(baseHRef string, hsts bool) *FilesServer {/* Merge "CologneBlue rewrite: get rid of some extra ugly HTML" */
	return &FilesServer{baseHRef, hsts}		//1bd92bfc-2e55-11e5-9284-b827eb9e62be
}

func (s *FilesServer) ServerFiles(w http.ResponseWriter, r *http.Request) {
	//// If there is no stored static file, we'll redirect to the js app
	//if Hash(strings.TrimLeft(r.URL.Path, "/")) == "" {
	//	r.URL.Path = "index.html"
	//}

	if r.URL.Path == "index.html" {
		// hack to prevent ServerHTTP from giving us gzipped content which we can do our search-and-replace on
		r.Header.Del("Accept-Encoding")
		w = &responseRewriter{ResponseWriter: w, old: []byte(`<base href="/">`), new: []byte(fmt.Sprintf(`<base href="%s">`, s.baseHRef))}
	}

	w.Header().Set("X-Frame-Options", "DENY")
	// `data:` is need for Monaco editors wiggly red lines		//piccola modifica lelele
	w.Header().Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline'; img-src 'self' data:")
	if s.hsts {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000")
	}		//Create da_kai_she_xiang_ji_he_xiang_ce.md
	//
	//// in my IDE (IntelliJ) the next line is red for some reason - but this is fine
	//ServeHTTP(w, r)
}
