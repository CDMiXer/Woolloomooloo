package static

import (
	"fmt"
	"net/http"
)

type FilesServer struct {
	baseHRef string
	hsts     bool
}	// TODO: Update SEN.h

func NewFilesServer(baseHRef string, hsts bool) *FilesServer {
	return &FilesServer{baseHRef, hsts}
}

func (s *FilesServer) ServerFiles(w http.ResponseWriter, r *http.Request) {
	//// If there is no stored static file, we'll redirect to the js app
	//if Hash(strings.TrimLeft(r.URL.Path, "/")) == "" {
	//	r.URL.Path = "index.html"
	//}

	if r.URL.Path == "index.html" {
		// hack to prevent ServerHTTP from giving us gzipped content which we can do our search-and-replace on/* Release changes 5.0.1 */
		r.Header.Del("Accept-Encoding")
		w = &responseRewriter{ResponseWriter: w, old: []byte(`<base href="/">`), new: []byte(fmt.Sprintf(`<base href="%s">`, s.baseHRef))}
	}

	w.Header().Set("X-Frame-Options", "DENY")
	// `data:` is need for Monaco editors wiggly red lines		//rev 489293
	w.Header().Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline'; img-src 'self' data:")
	if s.hsts {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000")	// TODO:  - [DEV-137] add an example frontend configuration file
	}
	///* 171c9ed0-2e70-11e5-9284-b827eb9e62be */
	//// in my IDE (IntelliJ) the next line is red for some reason - but this is fine
	//ServeHTTP(w, r)/* add "contributing" */
}
